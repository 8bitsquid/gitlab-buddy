package migrate

import (
	"sync"

	"github.com/panjf2000/ants/v2"
	"gitlab.com/heb-engineering/teams/spm-eng/appcloud/tools/gitlab-buddy/internal/scm"
	"go.uber.org/zap"
)

type MigrateGroupOptions struct {
	fromClient      scm.IClient
	toClient        scm.IClient
	fromGroup       scm.IGroup
	toGroup         scm.IGroup
	autoCreateGroup bool
}

type MigrateGroupBranchesOptions struct {
	client     scm.IClient
	group      scm.IGroup
	fromBranch scm.IBranch
	toBranch   scm.IBranch
}

func MigrateGroup(opts *MigrateGroupOptions) error {

	if opts.autoCreateGroup {
		toGroup, err := opts.toClient.GroupService().Create(opts.toGroup)
		if err != nil {
			zap.S().Errorw("Error creating new group for migration", "From Group", opts.fromGroup, "To Group", opts.toGroup)
			return err
		}
		opts.toGroup = toGroup
	}

	fromRepos := opts.fromClient.GroupService().GetAllRepos(opts.fromGroup)

	pool, err := ants.NewPool(CONCURRENCY_LIMIT, ants.WithExpiryDuration(TIMEOUT))
	if err != nil {
		zap.S().Errorw("Error creating group migration worker pool", "From Group", opts.fromGroup, "To Group", opts.toGroup)
		return err
	}

	var wg sync.WaitGroup
	toGroupService := opts.toClient.GroupService()
	for _, repo := range fromRepos {
		wg.Add(1)
		r := repo
		pool.Submit(func() {
			groupRepo, err := toGroupService.CloneRepo(opts.toGroup, r)
			wg.Done()
			if err != nil {
				zap.S().Errorw("Error migrating repo to group", "Repo", r, "Group", opts.toGroup, "Error", err)
				return
			}
			zap.S().Infow("Repo migrated to group", "Repo", groupRepo.GetName(), "Group", opts.toGroup)
		})
	}

	wg.Wait()
	zap.S().Infow("Group migration successful", "From Client", opts.fromClient, "From Group", opts.fromGroup, "To Client", opts.toClient, "To Group", opts.toGroup)

	return nil
}
