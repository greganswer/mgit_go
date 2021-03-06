package cmd

import (
	"errors"
	"fmt"

	"github.com/greganswer/mgit_go/git"
	"github.com/spf13/cobra"
)

// branchCmd represents the branch command
var branchCmd = &cobra.Command{
	Use:   "branch <issueID>",
	Short: "Create a branch using issue ID and title",
	Long: `The new branch name is taken from the title of the issue found.
The new branch is created off of the --base-branch or the
default base branch.`,
	Example: `  mgit branch JIR-123
  mgit branch JIR-123 --base-branch develop`,
	Args: branchCmdArgs,
	Run: branchCmdRun,
}

func init() {
	rootCmd.AddCommand(branchCmd)
	branchCmd.Flags().StringVar(&baseBranch, "base-branch", "", "the base branch to perform this action on")
}

// branchCmdArgs handles the optional arguments for the branch command.
func branchCmdArgs(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("requires issueID argument")
	}
	issueID = args[0]
	return nil
}

// branchCmdRun executes the branch command.
func branchCmdRun(cmd *cobra.Command, args []string) {
	i := getIssueFromTracker(issueID)

	prompt := fmt.Sprintf("Create %s branch from %s branch", emphasis(i.BranchName()), emphasis(baseBranch))
	if Confirm(prompt) {
		FailOrOK(git.CreateBranch(i.BranchName()))
	} else {
		skip("Branch not created")
	}
}
