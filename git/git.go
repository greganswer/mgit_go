package git

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/greganswer/mgit_go/issues"
)

// TODO: REMOVE ME
func todo(message string) {
	fmt.Println(color.YellowString("TODO:"), fmt.Sprintf("Implement git.%s", message))
	fmt.Println()
}

// AddAll stages all changes.
func AddAll() error {
	todo("AddAll")
	return nil
}

// CreateBranch creates a new git branch.
func CreateBranch(name string) error {
	todo("CreateBranch")
	return nil
	// return exec.Command("git", "checkout", "-b", name).Run()
}

// Commit the changes.
func Commit(message string) error {
	todo("Commit")
	return nil
}

// CurrentBranch returns the current branch for this Git repo.
func CurrentBranch() (string, error) {
	out, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	return strings.Trim(string(out), "\n"), err
}

// DefaultBaseBranch returns the base branch for this Git repo. Returns empty string if error.
func DefaultBaseBranch() string {
	todo("DefaultBaseBranch")
	// TODO: Return empty string if error.
	return "fake-default-base-branch"
}

// InitRepo initializes a Git repo.
func InitRepo() error {
	todo("InitRepo")
	return nil
}

// PullRequest uses the current branch to create a pull request on GitHub.
func PullRequest(issue issues.Issue) error {
	todo("PullRequest")
	// TODO: Get assignee from Viper config

	// TODO: Extract to viper config
	issueTracker := "GitHub"

	// TODO: Extract to viper config
	template := `%s

# [%s ticket %s](%s)

# Screenshots

# Sample API Requests

# QA Steps


# Checklist
- [] Added tests
- [] Check for typos
- [] Updated CHANGELOG.md
- [] Updated internal/external documentation`

	_ = fmt.Sprintf(template, issue.Title(), issueTracker, issue.ID, issues.URL(issue.ID))
	return nil
}

// Push updates to the remote repo.
func Push(branchName string) error {
	todo("Push")
	return nil
}

// Rebase changes from baseBranch.
func Rebase(baseBranch string) error {
	todo("Rebase")
	return nil
}
