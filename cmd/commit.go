package cmd

import (
	"github.com/greganswer/mgit_go/git"
	"github.com/spf13/cobra"
)

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Create a commit and push to GitHub",
	Long: `
All of the un-staged files are added, committed and pushed to GitHub.
The commit message is extracted from the branch name if one is not supplied
using the --message option. This command does the following:

  - Add all file (if commit was created)
  - Commit the changes (if commit was created)
  - Push changes to remote (if commit was created)
`,
	Example: `  mgit commit --message 'Update different from title'`,
	Run:     commitCmdRun,
}

func init() {
	rootCmd.AddCommand(commitCmd)
	commitCmd.Flags().StringVar(&commitMessage, "message", "", "the commit message")
	commitCmd.Flags().StringVar(&issueID, "issueID", "", "the ID of the issue being worked on")
}

// commitCmdRun executes the commit command.
func commitCmdRun(cmd *cobra.Command, args []string) {
	currentBranch, err := git.CurrentBranch()
	FailIfError(err)

	if commitMessage == "" {
		commitMessage = getCommitMessage(currentBranch)
	}

	// Ask to create commit.
	if confirmCommit(commitMessage, currentBranch) {
		addAllFiles()
		commitChanges(commitMessage)
		pushChanges(currentBranch)
	} else {
		skip("Changes not committed")
	}
}
