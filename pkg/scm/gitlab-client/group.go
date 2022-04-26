package gitlabclient

import (
	"runtime"
	"sync"

	"github.com/panjf2000/ants"
	"github.com/xanzy/go-gitlab"
	"gitlab.com/heb-engineering/teams/spm-eng/appcloud/tools/gitlab-buddy/pkg/scm"
	"go.uber.org/zap"
)

type GroupService struct {
	gc *GitlabClient
}

func (gs *GroupService) Get(groupID interface{}) (scm.IGroup, error) {
	group, resp, err := gs.gc.client.Groups.GetGroup(groupID)
	if err != nil {
		zap.S().Errorw("HTTP response getting group", "http_status", resp.StatusCode)
		return nil, err
	}

	g := &scm.Group{
		ID:   group.ID,
		Name: group.Name,
	}

	return g, nil
}

func (gs *GroupService) Create(group scm.IGroup) (scm.IGroup, error) {
	createOpts := &gitlab.CreateGroupOptions{
		Name: gitlab.String(group.GetName()),
	}
	newGroup, resp, err := gs.gc.client.Groups.CreateGroup(createOpts)
	if err != nil {
		zap.S().Errorw("Error creating group", "group", group, "http_status", resp.StatusCode, "error", err)
		return nil, err
	}

	g := &scm.Group{
		Name: newGroup.Name,
		ID:   newGroup.ID,
	}

	zap.S().Infow("Group group created", "group", newGroup)
	return g, nil
}

func (gs *GroupService) CloneRepo(group scm.IGroup, repo scm.IRepository) (scm.IRepository, error) {

	// update existing repo object with new group id
	repo.SetGroupID(group.GetID())
	clone, err := gs.gc.repos.Clone(repo)
	if err != nil {
		return nil, err
	}

	return clone, nil
}

func (gs *GroupService) GetAllRepos(group scm.IGroup) []scm.IRepository {
	page := 1
	repoPageList := make(chan []*gitlab.Project)
	var repos []scm.IRepository
	// var bar *progressbar.ProgressBar
	var wg sync.WaitGroup

	go func() {
		for list := range repoPageList {
			for _, repo := range list {
				// bar.Add(1)
				repos = append(repos, NewProject(repo))
			}
		}
	}()

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
		return nil
	}

	defer pool.Release()

	firstPage, resp, err := gs.getRepoListPage(group.GetID(), page)
	if err != nil {
		zap.S().Errorw("Error getting repo list page from group", "group_id", group, "page", page, "http_status", resp.StatusCode, "error", err)
		return nil
	}

	// barLabel := fmt.Sprintf("Getting repos from group %v", group)
	// bar = progressbar.NewOptions(resp.TotalItems, progressbar.OptionSetDescription(barLabel))
	repoPageList <- firstPage

	numPages := resp.TotalPages
	pagePool, err := ants.NewPool(scm.PAGINATOIN_CONCURRENCY_LIMIT, ants.WithExpiryDuration(scm.TIMEOUT))
	if err != nil {
		zap.S().Errorw("Error initializing repo list page worker pool", "group_id", group, "page", page, "error", err)
		return nil
	}

	for page <= numPages {
		page++
		p := page
		wg.Add(1)
		err = pagePool.Submit(func() {
			nextPage, resp, err := gs.getRepoListPage(group.GetID(), p)
			if err != nil {
				zap.S().Errorw("Error getting repo list page from group", "group_id", group, "page", page, "http_status", resp.StatusCode, "error", err)
				return
			}

			repoPageList <- nextPage
			wg.Done()
		})
		if err != nil {
			zap.S().Errorw("Error adding task to page pool for group repo list", "group_id", group, "page", page, "error", err)
		}
	}

	wg.Wait()
	close(repoPageList)

	return repos
}

func (gs *GroupService) getRepoListPage(group interface{}, page int) ([]*gitlab.Project, *gitlab.Response, error) {

	opts := &gitlab.ListGroupProjectsOptions{
		ListOptions: gitlab.ListOptions{
			Page:    page,
			PerPage: scm.PAGINATOIN_PER_PAGE,
		},
		IncludeSubgroups: gitlab.Bool(true),
	}

	return gs.gc.client.Groups.ListGroupProjects(group, opts)
}
