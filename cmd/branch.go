package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/greganswer/mgit_go/git"
	"github.com/greganswer/mgit_go/issues"
	"github.com/manifoldco/promptui"
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
		CheckIfError(err)
		finished()

		// Ask to create branch.
		if baseBranch == "" {
			baseBranch = git.DefaultBaseBranch()
		}
		branchInfo := fmt.Sprintf("%s from %s branch", info(issue.BranchName()), info(baseBranch))
		prompt := promptui.Prompt{
			Label:     fmt.Sprintf("Create %s?", branchInfo),
			IsConfirm: true,
		}

		_, err = prompt.Run()
		if err != nil {
			os.Exit(0)
		}

		// Create the branch.
		fmt.Printf("Creating %s...\n", branchInfo)
		finished()
	},
}

func init() {
	rootCmd.AddCommand(branchCmd)
	branchCmd.Flags().StringVar(&baseBranch, "base-branch", "", "the base branch to perform this action on")
}
