package cmd

import (
	"fmt"

	"github.com/greganswer/mgit_go/git"
	"github.com/greganswer/mgit_go/issues"
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
		// Get issue from current branch.
		currentBranch, err := git.CurrentBranch()
		FailIfError(err)
		issue, err := issues.FromBranch(currentBranch)
		FailIfError(err)

		// TODO: Get default value to work in init() below and remove these lines.
		if baseBranch == "" {
			baseBranch = git.DefaultBaseBranch()
		}

		// Ask to create commit.
		commitMessage := issue.String()
		// TODO: append "\n\nCloses #{issue.ID}" if issue tracker is GitHub.
		fmt.Printf("The commit message will be \"%s\"\n", info(commitMessage))
		if Confirm(fmt.Sprintf("Commit all changes to %s", info(currentBranch))) {
			// Add all file.
			fmt.Println("Adding all files...")
			finished()

			// Commit the changes.
			fmt.Println("Committing files...")
			// TODO: Use commitMessage
			finished()
		} else {
			skip("Changes not committed")
		}

		// Ask to rebase changes on base branch.
		if Confirm(fmt.Sprintf("Update the %s branch and rebase", info(baseBranch))) {
			fmt.Printf("Rebasing off of %s...\n", baseBranch)
			finished()
		} else {
			skip("No rebase off %s", info(baseBranch))
		}

		// Push changes to remote.
		fmt.Println("Pushing changes to origin...")
		finished()

		// Create pull request.
		fmt.Println("Opening pull request...")
		git.PullRequest(issue)
		finished()
	},
}

func init() {
	rootCmd.AddCommand(pullRequestCmd)
	pullRequestCmd.Flags().StringVar(&baseBranch, "base-branch", git.DefaultBaseBranch(), "the base branch to perform this action on")
	pullRequestCmd.Flags().StringVar(&issueID, "issueID", "", "the ID of the issue being worked on")
	pullRequestCmd.Flags().StringVar(&commitMessage, "message", "", "the commit message")
}
