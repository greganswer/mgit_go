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
		fmt.Printf("Retrieving issue data from %s...\n", info(issues.URL(issueID)))
		issue, err := issues.FromTracker(issueID)
		FailOrOK(err)
		branchName := issue.BranchName()

		// Create branch.
		branchInfo := fmt.Sprintf("%s branch from %s branch", info(branchName), info(baseBranch))
		if Confirm(fmt.Sprintf("Create %s", branchInfo)) {
			err = git.CreateBranch(branchName)
			FailOrOK(err)
		} else {
			skip("Branch not created")
		}
	},
}

func init() {
	rootCmd.AddCommand(branchCmd)
	branchCmd.Flags().StringVar(&baseBranch, "base-branch", "", "the base branch to perform this action on")
}
