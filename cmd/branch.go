package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// branchCmd represents the branch command
var branchCmd = &cobra.Command{
	Use:   "branch <issueId>",
	Short: "Create a branch using issue ID and title",
	Long: `The new branch name is taken from the title of the issue found.
The new branch is created off of the --base-branch or the default base branch.

NOTE
    User confirmation is required before the branch is created.

EXAMPLES
    $ mgit branch JIR-123
    $ mgit branch JIR-123 --base-branch develop`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires issueId argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("branch called with: " + strings.Join(args, " "))
	},
}
var baseBranch string

func init() {
	rootCmd.AddCommand(branchCmd)

	branchCmd.Flags().StringVar(&baseBranch, "base-branch", "", "the base branch to perform this action on")
}