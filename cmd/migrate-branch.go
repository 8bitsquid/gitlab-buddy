package cmd

import (
	// "errors"
	"errors"
	"os"

	"github.com/spf13/cobra"
	"gitlab.com/heb-engineering/teams/spm-eng/appcloud/tools/gitlab-buddy/pkg/migrate"
	gitlabclient "gitlab.com/heb-engineering/teams/spm-eng/appcloud/tools/gitlab-buddy/pkg/scm/gitlab-client"
	"go.uber.org/zap"
)

var (
	host                string
	repo                string
	group               string
	setDefault          bool
	setProtectedDefault bool
	archiveOldBranch    bool
	keepOldBranch       bool
	omitMergeRequests   bool
)

// branchCmd represents the branch command
var branchCmd = &cobra.Command{
	Use: "branch <from_branch> <to_branch>",
	Example: `glb migrate branch master main --repo "valid-repo-id" --host myhost
glb migrate branch master main --group="valid-group-id" --host=myhost --archive-old-branch --set-protected-deafult`,
	Short: "Migrates an existing branch to a new branch. Can be used on a single repo or all repos in a group",
	Long: `
Migrates an existing branch to a new branch. Originally intented to migrate repositories from [master] to [main], 
but can migrate/archive any Gitlab repo branches. Supports migration of an individual repo or all repos in a gorup.
See [flags] below for details on options for handling branches during migration.

`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Cobra apparently has a bug that omits the usage text from subcommands
		// This forces it to print the help text when no args are passed to this subcommand
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
		if len(args) < 2 {
			return errors.New("Two branch are required for migration. An existing branch and a new branch name to be created and migrated to.")
		}

		// Process the command args and flags
		client, err := gitlabclient.NewClient(host)
		if err != nil {
			return err
		}

		if repo == "" && group == "" {
			return errors.New("A valid group or repo id must be given.")
		} else if repo != "" && group != "" {
			return errors.New("Group and Repo flags cannot both be present.")
		}

		// MigrateBranchOptions are used in both single repo and group branch migrations
		// Setting the repo level migrate options without a repo defined.
		repoOpts := migrate.MigrateBranchOptions{
			Client:                client,
			OldBranch:             args[0],
			NewBranch:             args[1],
			SetAsDefault:          setDefault,
			SetAsProtectedDefault: setProtectedDefault,
			ArchiveOldBranch:      archiveOldBranch,
			KeepOldBranch:         keepOldBranch,
			OmitMergeRequests:     omitMergeRequests,
		}

		if repo != "" {
			// Get the gitlab project
			gitlabProject, err := client.RepoService().Get(repo)
			if err != nil {
				return err
			}
			// Update the migrate options with the repo
			repoOpts.Repo = gitlabProject
			zap.S().Debugw("Repo Migrate", "repo_ops", repoOpts)
			// Run single repo branch migration
			_, err = migrate.MigrateBranch(repoOpts)
			if err != nil {
				return err
			}
		} else {
			// Get that gitlab group yall
			gitlabGroup, err := client.GroupService().Get(group)
			if err != nil {
				return err
			}

			// Embed repoOpts into group migration opts struct

			groupOpts := migrate.MigrateBranchesInGroupOptions{
				Migrate: repoOpts,
				Group:   gitlabGroup,
			}

			zap.S().Debugw("Group Migrate", "group_ops", groupOpts)
			// Run group branch migration
			err = migrate.MigrateBranchesInGroup(groupOpts)
			if err != nil {
				return err
			}
		}

		return nil
	},
}

func init() {
	migrateCmd.AddCommand(branchCmd)

	branchCmd.Flags().StringVar(&host, "host", "", "Name of a host configured in the 'gitlab-buddy.yml' config. (Required if multiple hosts defined in the config file).")
	branchCmd.Flags().StringVarP(&repo, "repo", "r", "", "The Gitlab numberic ID or URL-encoded path of the repo/project. (e.g., heb/sub/group/project) See Gitlab's 'Namespaced path encoding' docs for more detail")
	branchCmd.Flags().StringVarP(&group, "group", "g", "", "The Gitlab numberic ID or URL-encoded path of the group/project. (e.g., heb/sub/group/project) See Gitlab's 'Namespaced path encoding' docs for more detail")
	branchCmd.Flags().BoolVarP(&setDefault, "set-default", "d", false, "Set the new branch as the default branch for the repo")
	branchCmd.Flags().BoolVarP(&setProtectedDefault, "set-protected-default", "p", true, "Set the new branch as default and protect it. This overrides '--set-default' flag.")
	branchCmd.Flags().BoolVarP(&archiveOldBranch, "archive-old-branch", "a", true, "Archive the old branch under the protected 'archive' tag and delete the old branch. The 'archive' tag is created/protected if not already in repo.")
	branchCmd.Flags().BoolVarP(&keepOldBranch, "keep-old-branch", "k", false, "Prevents old branch from being removed, as the old branch is removed by default")
	branchCmd.Flags().BoolVarP(&omitMergeRequests, "omit-merge-requests", "o", false, "Prevents open merge requests from being migrated - targeting the new branch.")

}
