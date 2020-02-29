package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// pullRequestCmd represents the pullRequest command
var pullRequestCmd = &cobra.Command{
	Use:     "pr",
	Aliases: []string{"pull-request"},
	Short:   "Create a GitHub Pull Request for the specified branch",
	Example: `    mgit pr
    mgit pr --base-branch develop`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Adding all files...")
		finished()

		fmt.Println("Commiting files...")
		finished()

		fmt.Println("Pushing changes to origin...")
		finished()

		fmt.Println("Opening pull request...")
		finished()
	},
}

func init() {
	rootCmd.AddCommand(pullRequestCmd)
	pullRequestCmd.Flags().StringVar(&baseBranch, "base-branch", "", "the base branch to perform this action on")
}
