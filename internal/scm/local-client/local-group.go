package localclient

import (
	"gitlab.com/heb-engineering/teams/spm-eng/appcloud/tools/gitlab-buddy/internal/scm"
	"go.uber.org/zap"
)

type Group struct {
	scm.Group
}

func NewGitGroup(name string) (scm.IGroup, error) {
	zap.S().Warn("Groups for local git not yet supported")
	return nil, nil
}

type GroupService struct {
	git *LocalClient
}

func (gr *GroupService) Get(path interface{}) (scm.IGroup, error) {
	zap.S().Warn("Groups for local git not yet supported")
	return nil, nil
}

func (gr *GroupService) Create(group scm.IGroup) (scm.IGroup, error) {
	zap.S().Warn("Groups for local git not yet supported")
	return nil, nil
}

func (gr *GroupService) CloneRepo(group scm.IGroup, repo scm.IRepository) (*scm.IRepository, error) {
	zap.S().Warn("Groups for local git not yet supported")
	return nil, nil
}

func (gr *GroupService) GetAllRepos(group scm.IGroup) []scm.IRepository {
	zap.S().Warn("Groups for local git not yet supported")
	return []scm.IRepository{}
}
