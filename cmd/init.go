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
  - Add all file (if user confirmed commit)
  - Commit the changes (if user confirmed commit)
  - Push changes to remote (if user confirmed commit)
`,
	Run: initCmdRun,
}

func init() {
	rootCmd.AddCommand(initCmd)
}

// initCmdRun executes the init command.
func initCmdRun(cmd *cobra.Command, args []string) {
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
		FailOrOK(git.InitRepo())
	}

	currentBranch, err := git.CurrentBranch()
	FailIfError(err)

	// Ask to create commit.
	if confirmCommit("Initial commit", currentBranch) {
		addAllFiles()
		commitChanges(commitMessage)
		pushChanges(currentBranch)
	} else {
		skip("Changes not committed")
	}
}
