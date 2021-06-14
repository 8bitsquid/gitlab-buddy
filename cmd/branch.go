package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	host                string
	repo                interface{}
	group               interface{}
	setDefault          bool
	setProtectedDefault bool
	archiveOldBranch    bool
	keepOldBranch       bool
	omitMergeRequests   bool
)

// branchCmd represents the branch command
var branchCmd = &cobra.Command{
	Use: "branch <from_branch> <to_branch>",
	Example: `glb migrate branch master main --repo="valid-repo-id" --host=myhost --archive-old-branch --set-protected-default
glb migrate branch master main --group="valid-group-id" --host=myhost --archive-old-branch --set-protected-deafult`,
	Short: "Migrates an existing branch to a new branch. Can be used on a single repo or all repos in a group",
	Long: `
Migrates a branch to a new branch. Originally intented to migrate repositories from [master] to [main], 
but can migrate/archive any Gitlab repo branches. Supports migration of an individual repo or all repos in a gorup.
See [flags] below for details on options for handling branches during migration.

`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		// zap.S().Debugw("cobra oto viper bind", "settings", viper.AllSettings())
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

func init() {
	migrateCmd.AddCommand(branchCmd)

	branchCmd.Flags().String("host", "", "Name of a host configured in the 'gitlab-buddy.yml' config")
	branchCmd.Flags().StringP("repo", "r", "", "The Gitlab numberic ID or URL-encoded path of the repo/project. (e.g., heb/sub/group/project) See Gitlab's 'Namespaced path encoding' docs for more detail")
	branchCmd.Flags().BoolP("set-default", "d", false, "Set the new branch as the default branch for the repo")
	branchCmd.Flags().BoolP("set-protected-default", "p", true, "Set the new branch as default and protect it. This overrides '--set-default' flag.")
	branchCmd.Flags().BoolP("archive-old-branch", "a", true, "Archive the old branch under the protected 'archive' tag and delete the old branch. The 'archive' tag is created/protected if not already in repo.")
	branchCmd.Flags().BoolP("keep-old-branch", "k", false, "Prevents old branch from being removed, as the old branch is removed by default")
	branchCmd.Flags().BoolP("omit-merge-requests", "o", false, "Prevents open merge requests from being migrated - targeting the new branch.")

}
