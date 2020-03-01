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
		issue, err := issues.FromBranch(git.CurrentBranch())
		FailIfError(err)

		fmt.Printf("The commit message will be \"%s\"\n", info(issue.String()))
		ConfirmOrAbort(fmt.Sprintf("Commit all changes to %s", info(git.CurrentBranch())))

		fmt.Println("Adding all files...")
		finished()

		fmt.Println("Committing files...")
		finished()

		fmt.Println("Pushing changes to origin...")
		finished()
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)
	commitCmd.Flags().StringVar(&commitMessage, "message", "", "the commit message")
	commitCmd.Flags().StringVar(&issueID, "issueID", "", "the ID of the issue being worked on")
}
