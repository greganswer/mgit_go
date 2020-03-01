package cmd

import (
	"errors"
	"fmt"

	"github.com/greganswer/mgit_go/git"
	"github.com/greganswer/mgit_go/issues"
	"github.com/spf13/cobra"
)

// branchCmd represents the branch command
var branchCmd = &cobra.Command{
	Use:   "branch <issueID>",
	Short: "Create a branch using issue ID and title",
	Long: `The new branch name is taken from the title of the issue found.
The new branch is created off of the --base-branch or the
default base branch.`,
	Example: `    mgit branch JIR-123
    mgit branch JIR-123 --base-branch develop`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires issueID argument")
		}
		issueID = args[0]
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Retrieve issue data.
		fmt.Printf("Retrieving issue data for %s...\n", info(issueID))
		issue, err := issues.FromTracker(issueID)
		FailIfError(err)
		finished()

		// Ask to create branch.
		// TODO: Extract method.
		if baseBranch == "" {
			baseBranch = git.DefaultBaseBranch()
		}
		branchInfo := fmt.Sprintf("%s branch from %s branch", info(issue.BranchName()), info(baseBranch))
		ConfirmOrAbort(fmt.Sprintf("Create %s", branchInfo))

		// Create the branch.
		fmt.Printf("Creating %s...\n", branchInfo)
		finished()
	},
}

func init() {
	rootCmd.AddCommand(branchCmd)
	branchCmd.Flags().StringVar(&baseBranch, "base-branch", "", "the base branch to perform this action on")
}
