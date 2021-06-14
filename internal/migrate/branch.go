package migrate

import (
	"sync"

	"gitlab.com/heb-engineering/teams/spm-eng/appcloud/tools/gitlab-buddy/internal/scm"
	"go.uber.org/zap"
)

const ARCHIVE_REPO_TAG = "archive"

type MigrateBranchOptions struct {
	Client                scm.IClient
	Repo                  scm.IRepository
	OldBranch             string
	NewBranch             string
	SetAsDefault          bool
	SetAsProtectedDefault bool
	ArchiveOldBranch      bool
	KeepOldBranch         bool
	OmitMergeRequests     bool
	MigrateSubmodules     bool
}

type MigrateBranchesInGroupOptions struct {
	Migrate MigrateBranchOptions
	Group   scm.IGroup
}

func MigrateBranch(opts MigrateBranchOptions) (scm.IBranch, error) {
	zap.S().Debugw("Migrating branch", "branch_migrate", opts)
	repos := opts.Client.RepoService()

	if !opts.MigrateSubmodules {
		if hasSubmod := repos.HasSubmodules(opts.Repo); hasSubmod {
			zap.S().Warnw("Skipping branch migration", "repo", opts.Repo.GetName(), "branch", opts.NewBranch)
			return nil, nil
		}
	}

	// Move branch
	branch, err := repos.MoveBranch(opts.Repo, opts.OldBranch, opts.NewBranch)
	if err != nil {
		zap.S().Errorw("Error moving branch", "Old Branch", opts.OldBranch, "New Branch", opts.NewBranch, "Repo", opts.Repo, "Error", err)
		return nil, err
	}

	// If repo.MoveBranch() returned nil, then references found to new branch found in repo files.
	// TODO: fix this - not explicite enough response from repo.MoveBranch()
	if branch == nil {
		zap.S().Debugw("Skipping: Unsafe to migrate branch", "migrate_branch", opts)
		return nil, nil
	}

	if opts.SetAsDefault || opts.SetAsProtectedDefault {
		branch, err = repos.SetDefaultBranch(opts.Repo, opts.NewBranch)
		if err != nil {
			zap.S().Errorw("Unable to set branch as default", "New Default", opts.NewBranch, "Repo", opts.Repo, "Error", err)
		}

		// If manually setting a protected default branch,
		// check if protect. If not protect that there branch .....yup
		if opts.SetAsProtectedDefault && !branch.IsProtected() {
			branch, err := repos.ProtectBranch(opts.Repo, opts.NewBranch)
			if err != nil {
				zap.S().Errorw("New default branch not protected", "New Default", opts.NewBranch, "Repo", opts.Repo, "Error", err)
			}
			zap.S().Debugw("New default branch protected", "branch", branch)
		}
	}

	// Archived tags are always protected after creation
	if opts.ArchiveOldBranch {
		// Add archive tag
		_, err := repos.AddTag(opts.Repo, ARCHIVE_REPO_TAG, opts.OldBranch, "Archived during branch migration via gitlab-buddy")
		if err != nil {
			// If error adding tag, consider archiving a failure and return
			zap.S().Errorw("Archive Old Branch Failed: Unable to create tag", "Branch", opts.OldBranch, "Repo", opts.Repo, "Error", err)
			return nil, err
		}

		// protect archive tag
		protTag, err := repos.ProtectTag(opts.Repo, ARCHIVE_REPO_TAG)
		if err != nil {
			zap.S().Errorw("Error protecting tag", "Tag", ARCHIVE_REPO_TAG, "Repo", opts.Repo, "Error", err)
			return nil, err
		}
		zap.S().Infow("Tag Protected", "Tag", protTag, "Repo", opts.Repo.GetName())
	}

	if !opts.KeepOldBranch {
		resp, err := repos.DeleteBranch(opts.Repo, opts.OldBranch)
		if err != nil {
			// Consider errors deleting old branch a "soft failure" and continue to return new/moved branch details
			zap.S().Errorw("Failed to delete old branch", "branch", opts.OldBranch, "repo", opts.Repo, "error", err, "response", resp.GetBody())
		}
	}

	if !opts.OmitMergeRequests {
		err := repos.UpdateMergeRequestsToNewBranch(opts.Repo, opts.OldBranch, opts.NewBranch)
		if err != nil {
			zap.S().Errorw("Failed updating merge requests to new branch", "error", err, "migrate_branch", opts)
		}
	}

	return branch, nil
}

func MigrateBranchesInGroup(opts MigrateBranchesInGroupOptions) error {
	zap.S().Debugw("Migrating group branch", "migrate_branch", opts)
	groupRepos := opts.Migrate.Client.GroupService().GetAllRepos(opts.Group)

	pool, err := getMigratePool()
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	for _, repo := range groupRepos {
		wg.Add(1)
		r := repo
		mOpts := opts.Migrate
		mOpts.Repo = r
		pool.Submit(func() {
			branch, err := MigrateBranch(mOpts)
			wg.Done()
			if err != nil {
				zap.S().Errorw("Error migrating branch in group repo", "migrate_group_branch", mOpts)
				return
			}
			zap.S().Infow("Group repo branch migration successful", "repo", mOpts.Repo.GetName(), "old_branch", mOpts.OldBranch, "new_branch", branch, "Migration", mOpts)
		})
	}

	wg.Wait()

	// Removes temporary dirs/files
	opts.Migrate.Client.Cleanup()
	return nil
}