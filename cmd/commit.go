package cmd

import (
	"fmt"

	"github.com/greganswer/mgit_go/git"
	"github.com/greganswer/mgit_go/issues"
	"github.com/spf13/cobra"
)

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Create a commit and push to GitHub",
	Long: `
All of the un-staged files are added, committed and pushed to GitHub.
The commit message is extracted from the branch name if one is not supplied
using the --message option.`,
	Example: `    mgit commit
    mgit commit --message 'Update different from title'`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get issue from current branch.
		currentBranch, err := git.CurrentBranch()
		FailIfError(err)
		issue, err := issues.FromBranch(currentBranch)
		FailIfError(err)

		// Ask to create commit.
		fmt.Printf("The commit message will be \"%s\"\n", info(issue.String()))
		if Confirm(fmt.Sprintf("Commit all changes to %s", info(currentBranch))) {
			// Add all file.
			fmt.Println("Adding all files...")
			err = git.AddAll()
			FailOrOK(err)

			// Commit the changes.
			fmt.Println("Committing files...")
			err = git.Commit("Initial commit")
			FailOrOK(err)

			// Push changes to remote.
			fmt.Printf("Pushing changes on %s branch to remote...\n", info(currentBranch))
			err = git.Push(currentBranch)
			FailOrOK(err)
		} else {
			skip("Changes not committed")
		}
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)
	commitCmd.Flags().StringVar(&commitMessage, "message", "", "the commit message")
	commitCmd.Flags().StringVar(&issueID, "issueID", "", "the ID of the issue being worked on")
}
