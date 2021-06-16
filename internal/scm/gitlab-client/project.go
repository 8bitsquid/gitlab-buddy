package gitlabclient

import (
	"net/http"
	"runtime"
	"sync"

	"github.com/panjf2000/ants"
	"github.com/xanzy/go-gitlab"
	"gitlab.com/heb-engineering/teams/spm-eng/appcloud/tools/gitlab-buddy/internal/scm"
	"gitlab.com/heb-engineering/teams/spm-eng/appcloud/tools/gitlab-buddy/tools"
	"go.uber.org/zap"
)

type Project struct {
	scm.Repository
	details *gitlab.Project
}

func NewProject(proj *gitlab.Project) scm.IRepository {
	return &Project{
		scm.Repository{
			ID:       proj.ID,
			Name:     proj.Name,
			Path:     proj.Path,
			URL:      proj.HTTPURLToRepo,
			CloneURL: proj.SSHURLToRepo,
		},
		proj,
	}
}

type ProjectService struct {
	gc *GitlabClient
}

func (ps *ProjectService) Get(projectID interface{}) (scm.IRepository, error) {
	repo, resp, err := ps.gc.client.Projects.GetProject(projectID, &gitlab.GetProjectOptions{})
	if err != nil {
		zap.S().Errorw("Unable to get project", "project_id", projectID, "http_status", resp.StatusCode)
		return nil, err
	}

	r := &scm.Repository{
		ID:       repo.ID,
		Name:     repo.Name,
		Path:     repo.Path,
		URL:      repo.WebURL,
		CloneURL: repo.SSHURLToRepo,
	}

	return r, nil
}

func (ps *ProjectService) Clone(repo scm.IRepository) (scm.IRepository, error) {
	project, err := ps.createProject(repo)
	if err != nil {
		return nil, err
	}

	// Set clone url of newly created project, from the seed repo
	project.SetCloneURL(repo.GetCloneURL())
	_, err = ps.gc.git.RepoService().Clone(project)
	if err != nil {
		return nil, err
	}

	return project, nil
}

func (ps *ProjectService) Push(repo scm.IRepository) error {
	err := ps.gc.git.RepoService().Push(repo)
	if err != nil {
		return err
	}

	return nil
}

func (ps *ProjectService) AddTag(repo scm.IRepository, tagName string, branch string, message string) (scm.ITag, error) {
	createOpts := &gitlab.CreateTagOptions{
		TagName: &tagName,
		Ref:     &branch,
		Message: &message,
	}

	// Check if tag already exists
	existingTag, resp, _ := ps.gc.client.Tags.GetTag(repo.GetID(), tagName)
	if resp.StatusCode < 400 {
		zap.S().Warnw("Tag already exists", "tag", tagName, "branch", branch, "repo", repo.GetName())
		return &scm.Tag{
			Name:   existingTag.Name,
			Commit: existingTag.Commit.ID,
		}, nil
	}

	// Create tag is doesn't exist
	tag, resp, err := ps.gc.client.Tags.CreateTag(repo.GetID(), createOpts)
	if err != nil {
		zap.S().Errorw("Error creating repo tag", "tag", tagName, "repo", repo, "http_status", resp.StatusCode)
		return nil, err
	}

	t := &scm.Tag{
		Name:   tag.Name,
		Commit: tag.Commit.ID,
	}

	zap.S().Infow("Tag Created", "tag", tag, "repo", repo.GetName())

	return t, nil
}

func (ps *ProjectService) ProtectTag(repo scm.IRepository, tagName string) (scm.ITag, error) {

	// Check if tag is already protected. if so, return
	_, resp, err := ps.gc.client.ProtectedTags.GetProtectedTag(repo.GetID(), tagName)
	if err != nil {
		if resp.StatusCode != http.StatusNotFound {
			zap.S().Debugw("Error checking if tag is already protected", "http_status", resp.StatusCode)
			return nil, err
		}
	} else {
		// TODO: Gitlab returns access level from Protected Tags API. Maybe intgrate that data into scm.ITag for future features
		return &scm.Tag{Name: tagName, Protected: true}, nil
	}

	protectOpts := &gitlab.ProtectRepositoryTagsOptions{
		Name: &tagName,
	}

	protectedTag, resp, err := ps.gc.client.ProtectedTags.ProtectRepositoryTags(repo.GetID(), protectOpts)
	if err != nil {
		zap.S().Errorw("Error protecting tag", "tag", tagName, "repo", repo.GetName(), "http_status", resp.StatusCode)
		return nil, err
	}

	pt := &scm.Tag{
		Name:      protectedTag.Name,
		Protected: true,
	}
	zap.S().Infow("Tag protected", "tag", tagName, "repo", repo.GetName())
	return pt, nil
}

func (ps *ProjectService) GetBranch(repo scm.IRepository, branch string) (scm.IBranch, error) {
	b, resp, err := ps.gc.client.Branches.GetBranch(repo.GetID(), branch)
	if err != nil {
		zap.S().Errorw("Failed to get branch", "http_status", resp.StatusCode, "branch", branch, "repo", repo)
		return nil, err
	}

	return &scm.Branch{
		Name:      b.Name,
		Default:   b.Default,
		Protected: b.Protected,
	}, nil
}

func (ps *ProjectService) MoveBranch(repo scm.IRepository, oldBranch string, newBranch string) (scm.IBranch, error) {

	pid := repo.GetID()

	// TODO: Find better way to do this - gitlab search API is way too limited and doesn't support regex or conditional logic (despite docs saying so)
	// Search for references to the newBranch in repo blobs
	queryList := tools.StringSandwich(oldBranch, "\"", "'")
	queryOpts := &gitlab.SearchOptions{
		PerPage: 100,
	}
	for _, query := range queryList {
		blobs, resp, err := ps.gc.client.Search.BlobsByProject(pid, query, queryOpts)
		if err != nil {
			zap.S().Errorw("Error searching project", "http_status", resp.StatusCode, "error", err)
			continue
		}
		if len(blobs) > 0 {
			zap.S().Errorw("References found to old branch in repo files.", "old_branch", oldBranch, "new_branch", newBranch, "repo", repo)
			return nil, nil
		}
	}

	zap.S().Debug("Safe to move branch: No references to old branch found in repo files", "old_branch", oldBranch, "new_branch", newBranch, "repo", repo.GetName())

	// Check if branch already exists, return with existing branch if it does
	existingBranch, resp, _ := ps.gc.client.Branches.GetBranch(pid, newBranch)
	if resp.StatusCode < 400 {
		zap.S().Warnw("Branch already exists", "branch", newBranch, "repo", repo.GetName())
		return &scm.Branch{
			Name: existingBranch.Name,
		}, nil
	}

	// Attempt to create branch
	branchOpts := &gitlab.CreateBranchOptions{
		Branch: &newBranch,
		Ref:    &oldBranch,
	}
	branch, resp, err := ps.gc.client.Branches.CreateBranch(pid, branchOpts)
	if err != nil {
		zap.S().Errorw("Unabled to create branch", "http_status", resp.StatusCode, "branch", newBranch, "repo", repo.GetName())
		return nil, err
	}
	zap.S().Infow("Branch successfully moved", "old_branch", oldBranch, "new_branch", newBranch, "repo", repo.GetName())

	b := &scm.Branch{
		Name: branch.Name,
	}
	return b, nil

}

func (ps *ProjectService) SetDefaultBranch(repo scm.IRepository, branch string) (scm.IBranch, error) {

	pid := repo.GetID()

	// Check if branch is already default, return if it is
	existingBranch, err := ps.GetBranch(repo, branch)
	if err == nil && existingBranch.IsDefault() {
		zap.S().Warnw("Branch is already default", "branch", branch, "repo", repo.GetName())
		return existingBranch, nil
	}

	// Gitlab API doesn't support setting default branch at the branch level
	// so we need to update the project, then retrieve our newly defaulted branch.
	// This allows us to easily check if the new default branch is auto-protected
	editProjOpts := &gitlab.EditProjectOptions{
		DefaultBranch: gitlab.String(branch),
	}
	_, resp, err := ps.gc.client.Projects.EditProject(pid, editProjOpts)
	if err != nil {
		zap.S().Error("Error setting default branch", "http_status", resp.StatusCode, "branch", branch, "repo", repo.GetName())
		return nil, err
	}

	defaultedBranch, err := ps.GetBranch(repo, branch)
	if err != nil {
		return nil, err
	}
	zap.S().Infof("%v default branch set to %v", repo.GetName(), branch)
	return defaultedBranch, nil
}

func (ps *ProjectService) DeleteBranch(repo scm.IRepository, branch string) (scm.IResponse, error) {
	zap.S().Infow("Attempting to delete branch", "branch", branch, "repo", repo.GetName())

	// check if branch is protected and attempt to unprotect
	protBranch, resp, err := ps.gc.client.Branches.GetBranch(repo.GetID(), branch)
	if err != nil {
		// Fail silently if branch doesn't exist
		if resp.StatusCode == http.StatusNotFound {
			zap.S().Warnw("Branch does not exist", "branch", branch, "repo", repo.GetName())
			return nil, nil
		}
		zap.S().Errorw("Faild to get branch", "branch", branch, "repo", repo.GetName(), "status", resp.StatusCode)
		return scm.NewResponse(resp.Response), err
	}
	if protBranch.Protected {
		unprotBranch, err := ps.UnprotectBranch(repo, branch)
		if err != nil {
			zap.S().Errorw("Failed trying to unprotect branch", "branch", branch, "repo", repo.GetName())
			return nil, err
		}
		zap.S().Infow("Branch protection removed", "branch", unprotBranch, "repo", repo.GetName())
	}

	resp, err = ps.gc.client.Branches.DeleteBranch(repo.GetID(), branch)
	if err != nil {
		zap.S().Errorw("Failed deleting branch", "branch", branch, "repo", repo.GetName())
		return scm.NewResponse(resp.Response), err
	}

	return scm.NewResponse(resp.Response), nil
}

func (ps *ProjectService) ProtectBranch(repo scm.IRepository, branch string) (scm.IBranch, error) {
	zap.S().Infof("Attempting to protect %v branch in %v", branch, repo.GetName())
	protectedBranch, resp, err := ps.gc.client.Branches.ProtectBranch(repo.GetID(), branch, &gitlab.ProtectBranchOptions{})
	if err != nil {
		zap.S().Errorw("Failed to protecte branch", "http_status", resp.StatusCode, "branch", branch, "repo", repo)
		return nil, err
	}
	zap.S().Infow("Branch protected", "branch", protectedBranch.Name, "repo", repo.GetName())
	return &scm.Branch{
		Name: protectedBranch.Name,
	}, nil
}

func (ps *ProjectService) UnprotectBranch(repo scm.IRepository, branch string) (scm.IBranch, error) {
	zap.S().Infof("Attempting to remove protection on %v branch in %v", branch, repo.GetName())
	unprotBranch, resp, err := ps.gc.client.Branches.UnprotectBranch(repo.GetID(), branch)
	if err != nil {
		zap.S().Errorw("Failed to remove branch protection", "http_status", resp.StatusCode, "branch", branch, "repo", repo)
		return nil, err
	}
	zap.S().Infow("Branch protection removed", "branch", unprotBranch.Name, "repo", repo.GetName())
	return &scm.Branch{
		Name: unprotBranch.Name,
	}, nil
}

// TODO: Improve scm.IRepository so that `oldBranch` can safely be inferred from `repo` in methods like this
func (ps *ProjectService) UpdateMergeRequestsToNewBranch(repo scm.IRepository, oldBranch string, newBranch string) error {
	zap.S().Infof("Attempting to update merge requests' target branch from %v to %v in %v", oldBranch, newBranch, repo.GetName())
	// List all merge requests that are opened, targeting repo.Branch
	listMROpts := &gitlab.ListProjectMergeRequestsOptions{
		TargetBranch: gitlab.String(oldBranch),
		State:        gitlab.String(RESOURCE_STATE_OPENED),
	}
	zap.S().Debugw("Checking for merge requests with options", "list_merge_request_opts", listMROpts)
	mrs, resp, err := ps.gc.client.MergeRequests.ListProjectMergeRequests(repo.GetID(), listMROpts)
	if err != nil {
		zap.S().Errorw("Failed to list project merge requests", "http_status", resp.StatusCode, "resource_state", listMROpts.State, "repo", repo.GetName())
		return err
	}

	if len(mrs) < 1 {
		zap.S().Infof("No merge requests found targeting %v", oldBranch)
		return nil
	}
	zap.S().Infof("%v merge requests targeting %v found for %v", len(mrs), oldBranch, repo.GetName())

	// Update merge requests
	var wg sync.WaitGroup
	// TODO: move this into scm/client.go as a utility function... or something better than this
	var numAnts int
	numCPUs := runtime.NumCPU()
	if numCPUs < scm.CONCURRENCY_LIMIT {
		numAnts = numCPUs
	} else {
		numAnts = scm.CONCURRENCY_LIMIT
	}

	pool, err := ants.NewPool(numAnts)
	if err != nil {
		zap.S().Errorw("Unable to initialize worker pool", "num_workers", numAnts, "error", err)
		return err
	}
	defer pool.Release()

	updateMROpts := &gitlab.UpdateMergeRequestOptions{
		TargetBranch: &newBranch,
	}
	for _, mr := range mrs {
		wg.Add(1)
		mergeRequest := mr
		zap.S().Debugw("Updating Merge Request", "merge_request", mergeRequest, "new_branch", newBranch)
		pool.Submit(func() {
			wg.Done()
			updatedMR, resp, err := ps.gc.client.MergeRequests.UpdateMergeRequest(repo.GetID(), mergeRequest.IID, updateMROpts)
			if err != nil {
				zap.S().Errorw("Unable to update project merge request", "http_status", resp.StatusCode, "merge_request", mergeRequest.Title, "repo", repo)
				zap.S().Warnw("Contact merge request author and reviewers about failed update", "merge_request_update", "failed", "author", mr.Author, "reviewer", mr.Reviewers, "assingee", mr.Assignee)
				return
			}
			zap.S().Infow("Project merge request updated", "merge_request", updatedMR.Title, "repo", repo.GetPath())
		})
	}

	wg.Wait()
	return nil
}

func (ps *ProjectService) HasSubmodules(repo scm.IRepository) bool {
	// Gitlab API only supports updating Submodules. So we need to use a
	// local client to clone repos into temp dirs to check for submodules
	// temp directories  should be removed with IRepository.Clean()
	localRepoService := ps.gc.git.RepoService()
	localClone, err := localRepoService.Clone(repo)
	if err != nil {
		zap.S().Errorw("Unable to clone or retrive repo for submodule check", "repo", repo.GetName(), "error", err)
	}

	if exists := localRepoService.HasSubmodules(localClone); exists {
		zap.S().Infow("Repo has submodules", "repo", repo.GetName())
		return true
	}

	return false
}

// Private methods
func (ps *ProjectService) createProject(repo scm.IRepository) (scm.IRepository, error) {
	proj := repo.(*Project)
	createOpts := ps.getCreatePayload(proj)
	zap.S().Debugw("Create Payload Recieved", "payload", createOpts)

	createdProj, resp, err := ps.gc.client.Projects.CreateProject(createOpts)
	if err != nil {
		zap.S().Errorw("Error creating repo", "http_status", resp.StatusCode, "repo", repo.GetName(), "group_id", repo.GetGroupID(), "host", ps.gc.client.BaseURL())
		return nil, err
	}

	zap.S().Infow("Repo created", "repo", createdProj.PathWithNamespace, "host", ps.gc.client.BaseURL())

	return NewProject(createdProj), nil
}

// TODO: Make generic reflective functions for converting Gitlab API schema
// This is dumb. Deep copy libraries tried had issues with reflection due to discrepencies between
// Gitlab's API scheme for get and crete project payloads
func (ps *ProjectService) getCreatePayload(p *Project) *gitlab.CreateProjectOptions {
	payload := &gitlab.CreateProjectOptions{
		Name:                             &p.details.Name,
		Path:                             &p.details.Path,
		NamespaceID:                      &p.GroupID,
		DefaultBranch:                    &p.details.DefaultBranch,
		Description:                      &p.details.Description,
		IssuesAccessLevel:                &p.details.IssuesAccessLevel,
		RepositoryAccessLevel:            &p.details.RepositoryAccessLevel,
		MergeRequestsAccessLevel:         &p.details.MergeRequestsAccessLevel,
		ForkingAccessLevel:               &p.details.ForkingAccessLevel,
		BuildsAccessLevel:                &p.details.BuildsAccessLevel,
		WikiAccessLevel:                  &p.details.WikiAccessLevel,
		SnippetsAccessLevel:              &p.details.SnippetsAccessLevel,
		PagesAccessLevel:                 &p.details.PagesAccessLevel,
		OperationsAccessLevel:            &p.details.OperationsAccessLevel,
		ResolveOutdatedDiffDiscussions:   &p.details.ResolveOutdatedDiffDiscussions,
		ContainerRegistryEnabled:         &p.details.ContainerRegistryEnabled,
		SharedRunnersEnabled:             &p.details.SharedRunnersEnabled,
		Visibility:                       &p.details.Visibility,
		PublicBuilds:                     &p.details.PublicBuilds,
		AllowMergeOnSkippedPipeline:      &p.details.AllowMergeOnSkippedPipeline,
		OnlyAllowMergeIfPipelineSucceeds: &p.details.OnlyAllowMergeIfPipelineSucceeds,
		OnlyAllowMergeIfAllDiscussionsAreResolved: &p.details.OnlyAllowMergeIfAllDiscussionsAreResolved,
		MergeMethod:                     &p.details.MergeMethod,
		RemoveSourceBranchAfterMerge:    &p.details.RemoveSourceBranchAfterMerge,
		LFSEnabled:                      &p.details.LFSEnabled,
		RequestAccessEnabled:            &p.details.RequestAccessEnabled,
		TagList:                         &p.details.TagList,
		PrintingMergeRequestLinkEnabled: &p.details.MergeRequestsEnabled,
		BuildCoverageRegex:              &p.details.BuildCoverageRegex,
		CIConfigPath:                    &p.details.CIConfigPath,
		CIForwardDeploymentEnabled:      &p.details.CIForwardDeploymentEnabled,
		ApprovalsBeforeMerge:            &p.details.ApprovalsBeforeMerge,
		Mirror:                          &p.details.Mirror,
		MirrorTriggerBuilds:             &p.details.MirrorTriggerBuilds,
		PackagesEnabled:                 &p.details.PackagesEnabled,
		ServiceDeskEnabled:              &p.details.ServiceDeskEnabled,
		AutocloseReferencedIssues:       &p.details.AutocloseReferencedIssues,
		SuggestionCommitMessage:         &p.details.SuggestionCommitMessage,
		IssuesEnabled:                   &p.details.IssuesEnabled,
		MergeRequestsEnabled:            &p.details.MergeRequestsEnabled,
		JobsEnabled:                     &p.details.JobsEnabled,
		WikiEnabled:                     &p.details.WikiEnabled,
		SnippetsEnabled:                 &p.details.SnippetsEnabled,
	}

	return payload
}
