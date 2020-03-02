package cmd

import (
	"fmt"

	"github.com/greganswer/mgit_go/file"
	"github.com/greganswer/mgit_go/git"
	"github.com/greganswer/mgit_go/issues"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize local repository and push to remote",
	Run: func(cmd *cobra.Command, args []string) {
		// Create .gitignore file.
		fmt.Println("Creating .gitignore file...")
		exists, err := file.Touch(".gitignore")
		FailIfError(err)
		if exists {
			skip(".gitignore file already exists")
		} else {
			finished()
		}

		// Initialize a git repo.
		fmt.Println("Initializing repo...")
		exists, err = file.Exists(".git")
		FailIfError(err)
		if exists {
			skip("Repo already intialized")
		} else {
			// TODO: Initialize repo
			finished()
		}

		// Ask to create commit.
		issue, err := issues.FromBranch(git.CurrentBranch())
		FailIfError(err)

		fmt.Printf("The commit message will be \"%s\"\n", info(issue.String()))
		if Confirm(fmt.Sprintf("Commit all changes to %s", info(git.CurrentBranch()))) {
			// Add all file.
			fmt.Println("Adding all files...")
			finished()

			// Commit the changes.
			fmt.Println("Committing files...")
			finished()
		} else {
			skip("Changes not committed")
		}

		// Push changes to remote.
		fmt.Println("Pushing changes to origin...")
		finished()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
