package cmd

import (
	"fmt"

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
		fmt.Println("Adding all files...")
		finished()

		fmt.Println("Commiting files...")
		finished()

		fmt.Println("Pushing changes to origin...")
		finished()
	},
}

func init() {
	defaultMessage := "Initial commit" // TODO: Update this with title from issue.
	defaultIssueID := "JIR-123"        // TODO: Update this with ID from issue.

	rootCmd.AddCommand(commitCmd)
	commitCmd.Flags().StringVar(&commitMessage, "message", defaultMessage, "the commit message")
	commitCmd.Flags().StringVar(&issueID, "issueID", defaultIssueID, "the ID of the issue being worked on")
}
