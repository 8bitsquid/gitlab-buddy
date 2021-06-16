package localclient

import (
	"errors"
	"fmt"
	"path/filepath"

	"gitlab.com/heb-engineering/teams/spm-eng/appcloud/tools/gitlab-buddy/internal/config"
	"gitlab.com/heb-engineering/teams/spm-eng/appcloud/tools/gitlab-buddy/internal/scm"
	"go.uber.org/zap"
)

type Repo struct {
	scm.Repository
}

func NewRepo(path string) scm.IRepository {
	return &Repo{
		scm.Repository{
			Path: path,
		},
	}
}

type RepoService struct {
	git *LocalClient
}

func (gr *RepoService) Get(path interface{}) (scm.IRepository, error) {
	// check if dir exists
	exists, err := config.ExistsOnDisk(path.(string))
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, errors.Unwrap(fmt.Errorf("local path not found: %w", path))
	}

	r := &Repo{
		scm.Repository{
			Path: path.(string),
		},
	}
	return r, nil
}

func (gr *RepoService) Clone(repo scm.IRepository) (scm.IRepository, error) {
	// Check if local clone already exists and return if it does

	checkPath := filepath.Join(gr.git.baseDir, repo.GetPath())
	if exists, _ := config.ExistsOnDisk(checkPath); exists {
		return repo, nil
	}

	resp, err := gr.git.cmd.FromBaseDir(gr.git.baseDir).Clone().Repo(repo.GetCloneURL()).Exec()
	if err != nil {
		zap.S().Errorw("Error cloning repo", "Repo", repo, "Response", resp, "error", err)
		return nil, err
	}

	zap.S().Debugw("Clone.Repo Command Response", "clone_repo_cmd_response", resp)
	clone := &Repo{
		scm.Repository{
			Path: filepath.Join(gr.git.baseDir, repo.GetPath()),
			Name: repo.GetName(),
		},
	}
	return clone, nil
}

func (gr *RepoService) Push(repo scm.IRepository) error {
	_, err := gr.git.cmd.FromBaseDir(repo.GetPath()).Push().Repo(repo.GetName()).Exec()
	return err
}

func (rs *RepoService) AddTag(repo scm.IRepository, tagName string, commit string, message string) (scm.ITag, error) {
	if rs.tagExists(repo, tagName) {
		return nil, errors.Unwrap(fmt.Errorf("Tag %w already exists in repo %w", tagName, repo.GetPath()))
	}

	_, err := rs.git.cmd.FromBaseDir(repo.GetPath()).FromBaseDir(repo.GetPath()).Tag().CreateTag(tagName).WithBranch(commit).WithMessage(message).Exec()
	if err != nil {
		return nil, err
	}

	return &scm.Tag{
		Name:    tagName,
		Commit:  commit,
		Message: message,
	}, nil
}

func (rs *RepoService) ProtectTag(repo scm.IRepository, tagName string) (scm.ITag, error) {
	zap.S().Warn("Protect tags are not supported for local git repos")
	return nil, nil
}

// TODO: actually get branch
func (rs *RepoService) GetBranch(repo scm.IRepository, branch string) (scm.IBranch, error) {
	return nil, nil
}

func (rs *RepoService) MoveBranch(repo scm.IRepository, oldBranch string, newBranch string) (scm.IBranch, error) {
	_, err := rs.git.cmd.FromBaseDir(repo.GetPath()).Branch().Move(oldBranch).To(newBranch).Exec()
	if err != nil {
		return nil, err
	}

	return &scm.Branch{
		Name: newBranch,
	}, nil
}

func (rs *RepoService) SetDefaultBranch(repo scm.IRepository, newDefault string) (scm.IBranch, error) {
	zap.S().Error("Setting default branch (i.e., HEAD) for local git not supported")
	return nil, nil
}

func (rs *RepoService) DeleteBranch(repo scm.IRepository, branch string) (scm.IResponse, error) {
	resp, err := rs.git.cmd.FromBaseDir(repo.GetPath()).Branch().Delete(branch).Exec()
	if err != nil {
		return &scm.Response{BodyString: resp}, err
	}
	return &scm.Response{BodyString: resp}, err
}

func (gr *RepoService) ProtectBranch(repo scm.IRepository, branch string) (scm.IBranch, error) {
	zap.S().Error("Branch protection for local git not supported")
	return nil, nil
}

func (gr *RepoService) UnprotectBranch(repo scm.IRepository, branch string) (scm.IBranch, error) {
	zap.S().Error("Branch protection for local git not supported")
	return nil, nil
}

func (rs *RepoService) UpdateMergeRequestsToNewBranch(repo scm.IRepository, oldBranch string, newBranch string) error {
	zap.S().Error("Merge requests for local git not supported")
	return nil
}

func (rs *RepoService) HasSubmodules(repo scm.IRepository) bool {
	path := repo.GetPath()
	resp, err := rs.git.cmd.FromBaseDir(path).Submodule().Status().Exec()
	if err != nil {
		zap.S().Errorw("Failed checking for Sumodules", "error", err)
	}
	return resp != ""
}

func (rs *RepoService) tagExists(repo scm.IRepository, tagName string) bool {
	tag, err := rs.git.cmd.FromBaseDir(repo.GetPath()).Tag().GetTag(tagName).Exec()
	if err != nil {
		zap.S().Errorw("Error check if tag exists in repo", "tag", tagName, "repo", repo, "error", err)
		return false
	}
	return tag != ""
}
