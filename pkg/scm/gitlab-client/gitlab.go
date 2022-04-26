package gitlabclient

import (
	"github.com/xanzy/go-gitlab"
	"go.uber.org/zap"

	"gitlab.com/heb-engineering/teams/spm-eng/appcloud/tools/gitlab-buddy/pkg/config"
	"gitlab.com/heb-engineering/teams/spm-eng/appcloud/tools/gitlab-buddy/pkg/scm"
	localclient "gitlab.com/heb-engineering/teams/spm-eng/appcloud/tools/gitlab-buddy/pkg/scm/local-client"
)

const (
	TEMP_DIR_SUB_PATH = "gitlabbuddy_gc"

	RESOURCE_STATE_OPENED = "opened"
)

type GitlabClient struct {
	client *gitlab.Client
	git    *localclient.LocalClient
	repos  scm.IRepoService
	groups scm.IGroupService
}

func NewClient(name string) (scm.IClient, error) {
	host, err := config.GetHost(name)
	if err != nil {
		return nil, err
	}
	baseURL := host.GetBaseURL()

	gitlabClient, err := gitlab.NewClient(host.Token, gitlab.WithBaseURL(baseURL.String()))
	if err != nil {
		return nil, err
	}

	// Check that the user for the token exists
	user, resp, err := gitlabClient.Users.CurrentUser()
	if err != nil {
		if resp.StatusCode == 401 {
			zap.S().Errorw("User unauthorized. Ensure the API key is valid and has proper permissions.", "host", host)
		}
		return nil, err
	}

	if !user.IsAdmin {
		zap.S().Warn("Current Gitlab user is not an admin. API methods may be limited.")
	}

	git, err := localclient.NewGitClientWithTempDir(TEMP_DIR_SUB_PATH)
	if err != nil {
		return nil, err
	}

	gc := &GitlabClient{
		client: gitlabClient,
		git:    git,
	}
	gc.repos = &ProjectService{gc: gc}
	gc.groups = &GroupService{gc: gc}

	zap.S().Debugw("New remote Gitlab client created", "client", gc)

	return gc, nil
}

func (g GitlabClient) GroupService() scm.IGroupService {
	return g.groups
}

func (g GitlabClient) RepoService() scm.IRepoService {
	return g.repos
}

func (g *GitlabClient) Cleanup() error {

	err := g.git.Cleanup()
	if err != nil {
		return err
	}

	return nil
}
