package cmd

import (
	"fmt"

	"github.com/greganswer/mgit_go/file"
	"github.com/greganswer/mgit_go/git"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize local repository and push to remote",
	Long: `This command is idempotent and it will prompt you before commiting changes and
pushing them to the remote repo. This command does the following:

    - Create .gitignore file (if not already created)
    - Initialize a git repo (if not already intialized)
    - Asks to create an initial commit
    - Add all file (if commit was created)
    - Commit the changes (if commit was created)
    - Push changes to remote (if commit was created)
`,
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
			err = git.InitRepo()
			FailOrOK(err)
		}

		// Ask to create commit.
		currentBranch, err := git.CurrentBranch()
		FailIfError(err)
		fmt.Printf("The commit message will be \"Initial commit\"\n")
		prompt := fmt.Sprintf("Commit all changes to the \"%s\" branch", info(currentBranch))
		if Confirm(prompt) {
			// Add all file.
			fmt.Println("Adding all files...")
			err = git.AddAll()
			FailOrOK(err)

			// Commit the changes.
			fmt.Println("Committing files...")
			err = git.Commit("Initial commit")
			FailOrOK(err)

			// Push changes to remote.
			fmt.Printf("Pushing changes on %s branch to remote...\n", info(currentBranch))
			err = git.Push(currentBranch)
			FailOrOK(err)
		} else {
			skip("Changes not committed")
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
