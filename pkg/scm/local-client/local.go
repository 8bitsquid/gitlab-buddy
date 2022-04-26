package localclient

import (
	"github.com/spf13/afero"
	"gitlab.com/heb-engineering/teams/spm-eng/appcloud/tools/gitlab-buddy/pkg/config"
	gitcmd "gitlab.com/heb-engineering/teams/spm-eng/appcloud/tools/gitlab-buddy/pkg/pkg/git-cmd"
	"gitlab.com/heb-engineering/teams/spm-eng/appcloud/tools/gitlab-buddy/pkg/scm"
	"go.uber.org/zap"
)

type LocalClient struct {
	baseDir         string
	isTempDirClient bool
	fs              afero.Fs // fs applied only to the same directory
	osfs            afero.Fs //the root os file system
	cmd             *gitcmd.GitCommand
	groupService    *GroupService
	repoService     *RepoService
}

func NewLocalClient(baseDir string) (*LocalClient, error) {
	baseExists, err := config.ExistsOnDisk(baseDir)
	if err != nil {
		return nil, err
	}

	var fs afero.Fs
	if baseExists {
		fs = afero.NewBasePathFs(afero.NewOsFs(), baseDir)
	}

	gitcmd := gitcmd.NewGitCommand()

	client := &LocalClient{
		baseDir: baseDir,
		fs:      fs,
		cmd:     gitcmd,
	}
	client.groupService = &GroupService{git: client}
	client.repoService = &RepoService{git: client}

	return client, nil
}

// TODO: Duplicated processes - improve NewGitClient() method with tempdir options
func NewGitClientWithTempDir(prefix string) (*LocalClient, error) {
	osfs := afero.NewOsFs()
	tmpDir, err := afero.TempDir(osfs, "", prefix)

	if err != nil {
		return nil, err
	}

	client := &LocalClient{
		baseDir:         tmpDir,
		isTempDirClient: true,
		fs:              afero.NewBasePathFs(osfs, tmpDir),
		osfs:            osfs,
		cmd:             gitcmd.NewGitCommand(),
	}
	client.groupService = &GroupService{git: client}
	client.repoService = &RepoService{git: client}

	return client, nil
}

func (gc *LocalClient) GroupService() scm.IGroupService {
	zap.S().Error("Local client groups not supported")
	return nil
}

func (gc *LocalClient) RepoService() scm.IRepoService {
	return gc.repoService
}

func (gc *LocalClient) Cleanup() error {
	if gc.isTempDirClient {
		zap.S().Debugw("Attempting to remove temporary dir", "dir", gc.baseDir)
		err := gc.osfs.RemoveAll(gc.baseDir)
		if err != nil {
			return err
		}
	}
	return nil
}
