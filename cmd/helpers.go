package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/greganswer/mgit_go/git"
	"github.com/greganswer/mgit_go/issues"
	"github.com/manifoldco/promptui"
)

// Fail exits the program with a standardized error message.
func Fail(err error) {
	red := color.New(color.FgRed, color.Bold).SprintFunc()
	fmt.Println(red("FAIL:"), err)
	os.Exit(1)
}

// FailIfError exits the program with a standardized error message if an error occurred.
func FailIfError(err error) {
	if err != nil {
		Fail(err)
	}
}

// FailOrOK displays a failure message if there's an error otherwise displays "OK".
func FailOrOK(err error) {
	FailIfError(err)
	finished()
}

// Confirm returns true if the user confirms or if the yes flag is present.
func Confirm(message string) bool {
	if yes {
		return true
	}
	prompt := promptui.Prompt{
		Label:     message,
		IsConfirm: true,
	}
	_, err := prompt.Run()
	return err == nil
}

func skip(messages ...interface{}) {
	message := fmt.Sprintf(messages[0].(string), messages[1:]...)
	fmt.Println(color.YellowString("SKIP:"), message)
	fmt.Println("")
}

func warn(messages ...interface{}) {
	message := fmt.Sprintf(messages[0].(string), messages[1:]...)
	fmt.Println(color.YellowString("WARN:"), message)
}

func emphasis(message string) string {
	return color.CyanString(message)
}

func finished() {
	c := color.New(color.FgGreen, color.Bold)
	c.Println("OK")
	fmt.Println()
}

// getCommitMessage get a commit message from the branch.
func getCommitMessage(branchName string) string {
	issue, err := issues.FromBranch(branchName)
	switch err.(type) {
	case *issues.BranchHasNoIDError:
		warn(err.Error())
	default:
		Fail(err)
	}
	return issue.String()
}

func addAllFiles() {
	fmt.Println("Adding all files...")
	FailOrOK(git.AddAll())
}

func commitChanges(message string) {
	fmt.Println("Committing files...")
	FailOrOK(git.Commit(message))
}

func pushChanges(branch string) {
	fmt.Printf("Pushing changes on the %s branch to remote...\n", emphasis(branch))
	FailOrOK(git.Push(branch))
}

func getIssueFromCurrentBranch() issues.Issue {
	currentBranch, err := git.CurrentBranch()
	FailIfError(err)
	i, err := issues.FromBranch(currentBranch)
	FailIfError(err)
	return i
}

func getIssueFromTracker(issueID string) issues.Issue {
	fmt.Printf("Retrieving issue data from %s...\n", emphasis(issues.APIURL(issueID)))
	i, err := issues.FromTracker(issueID)
	FailOrOK(err)
	return i
}

func confirmCommit(commitMessage, currentBranch string) bool {
	if commitMessage != "" {
		fmt.Printf("The commit message will be \"%s\"\n", emphasis(commitMessage))
	}

	prompt := fmt.Sprintf("Commit all changes to the %s branch", emphasis(currentBranch))
	return Confirm(prompt)
}