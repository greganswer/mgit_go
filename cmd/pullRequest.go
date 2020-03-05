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
	Long: `This command is idempotent and it will prompt you before commiting changes and
pushing them to the remote repo. This command does the following:

  - Asks to create a commit
  - Add all file (if user confirmed commit)
  - Commit the changes (if user confirmed commit)
  - Rebase changes off base branch (if user confirms rebase)
  - Push changes to remote 
  - Create the pull request on GitHub
`,
	Example: `  mgit pr --base-branch develop`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get issue from current branch.
		currentBranch, err := git.CurrentBranch()
		FailIfError(err)
		issue, err := issues.FromBranch(currentBranch)
		FailIfError(err)

		// Ask to create commit.
		if commitMessage == "" {
			commitMessage = getCommitMessage(currentBranch)
		}
		// TODO: append "\n\nCloses #{issue.ID}" if issue tracker is GitHub.
		fmt.Printf("The commit message will be \"%s\"\n", emphasis(commitMessage))
		if Confirm(fmt.Sprintf("Commit all changes to %s branch", emphasis(currentBranch))) {
			addAllFiles()
			commitChanges(commitMessage)
		} else {
			skip("Changes not committed")
		}

		// Rebase changes on base branch.
		prompt := fmt.Sprintf("Update the %s branch and rebase", emphasis(baseBranch))
		if Confirm(prompt) {
			fmt.Printf("Rebasing off of %s...\n", emphasis(baseBranch))
			FailOrOK(git.Rebase(baseBranch))
			pushChanges(currentBranch)
		} else {
			skip("No rebase off %s", emphasis(baseBranch))
		}

		pushChanges(currentBranch)

		// Create pull request.
		fmt.Println("Opening pull request...")
		FailOrOK(git.PullRequest(issue))
	},
}

func init() {
	rootCmd.AddCommand(pullRequestCmd)
	pullRequestCmd.Flags().StringVar(&baseBranch, "base-branch", "", "the base branch to perform this action on")
	pullRequestCmd.Flags().StringVar(&commitMessage, "message", "", "the commit message")
}
