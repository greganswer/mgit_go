package cmd

import (
	"fmt"

	"github.com/greganswer/mgit_go/git"
	"github.com/greganswer/mgit_go/issues"
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
)

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open an issue in the default web browser",
	Run: func(cmd *cobra.Command, args []string) {
		// Get issue from current branch.
		currentBranch, err := git.CurrentBranch()
		FailIfError(err)
		issue, err := issues.FromBranch(currentBranch)
		FailIfError(err)

		// Open the issue in web browser.
		fmt.Println("Opening issue...")
		err = browser.OpenURL(issues.URL(issue.ID))
		FailIfError(err)
	},
}

func init() {
	rootCmd.AddCommand(openCmd)
}
