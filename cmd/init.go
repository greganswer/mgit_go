package cmd

import (
	"fmt"

	"github.com/greganswer/mgit_go/git"
	"github.com/greganswer/mgit_go/issues"
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

		issue, err := issues.FromBranch(git.CurrentBranch())
		FailIfError(err)

		fmt.Printf("The commit message will be \"%s\"\n", info(issue.String()))
		ConfirmOrAbort(fmt.Sprintf("Commit all changes to %s", info(git.CurrentBranch())))

		fmt.Println("Adding all files...")
		finished()

		fmt.Println("Committing files...")
		finished()

		fmt.Println("Pushing changes to origin...")
		finished()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
