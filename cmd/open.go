package cmd

import (
	"fmt"

	"github.com/pkg/browser"
	"github.com/spf13/cobra"
)

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open an issue in the default web browser",
	Run: func(cmd *cobra.Command, args []string) {
		i := getIssueFromCurrentBranch()
		fmt.Println("Opening issue...")
		FailIfError(browser.OpenURL(i.WebURL()))
	},
}

func init() {
	rootCmd.AddCommand(openCmd)
}
