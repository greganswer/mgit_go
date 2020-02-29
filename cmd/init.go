package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize local repository and push to remote",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Creating .gitignore file...")
		finished()

		fmt.Println("Initializing repo...")
		finished()

		commitAndPushAllNewChanges()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
