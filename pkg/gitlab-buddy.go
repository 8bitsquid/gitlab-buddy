package pkg

import (
	"gitlab.com/heb-engineering/teams/spm-eng/appcloud/tools/gitlab-buddy/pkg/scm"
	gitlabclient "gitlab.com/heb-engineering/teams/spm-eng/appcloud/tools/gitlab-buddy/pkg/scm/gitlab-client"
	localclient "gitlab.com/heb-engineering/teams/spm-eng/appcloud/tools/gitlab-buddy/pkg/scm/local-client"
	"go.uber.org/zap"
)

func NewRemote(name string) (scm.IClient, error) {

	// TODO: Only gitlab remote hosts supported - update if others are added
	client, err := gitlabclient.NewClient(name)
	if err != nil {
		return nil, err
	}

	zap.S().Debugw("New Remote Client created", "client", client)

	return client, nil
}

func NewLocal(path string) (scm.IClient, error) {
	client, err := localclient.NewLocalClient(path)
	if err != nil {
		zap.S().Errorw("Unable to create local directory git client", "path", path)
	}

	return client, nil
}
