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
		commitMessage := issue.String()
		// TODO: append "\n\nCloses #{issue.ID}" if issue tracker is GitHub.
		fmt.Printf("The commit message will be \"%s\"\n", info(commitMessage))
		if Confirm(fmt.Sprintf("Commit all changes to %s", info(currentBranch))) {
			// Add all file.
			fmt.Println("Adding all files...")
			err = git.AddAll()
			FailOrOK(err)

			// Commit the changes.
			fmt.Println("Committing files...")
			err = git.Commit("Initial commit")
			FailOrOK(err)
		} else {
			skip("Changes not committed")
		}

		// Rebase changes on base branch.
		if Confirm(fmt.Sprintf("Update the %s branch and rebase", info(baseBranch))) {
			fmt.Printf("Rebasing off of %s...\n", info(baseBranch))
			err = git.Rebase(baseBranch)
			FailOrOK(err)
		} else {
			skip("No rebase off %s", info(baseBranch))
		}

		// Push changes to remote.
		fmt.Printf("Pushing changes on %s branch to remote...\n", info(currentBranch))
		err = git.Push(currentBranch)
		FailOrOK(err)

		// Create pull request.
		fmt.Println("Opening pull request...")
		err = git.PullRequest(issue)
		FailOrOK(err)
	},
}

func init() {
	rootCmd.AddCommand(pullRequestCmd)
	pullRequestCmd.Flags().StringVar(&baseBranch, "base-branch", "", "the base branch to perform this action on")
	pullRequestCmd.Flags().StringVar(&commitMessage, "message", "", "the commit message")
}
