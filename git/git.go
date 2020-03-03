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
}

// CurrentBranch returns the current branch for this Git repo.
func CurrentBranch() (string, error) {
	out, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	return strings.Trim(string(out), "\n"), err
}

// DefaultBaseBranch returns the base branch for this Git repo.
func DefaultBaseBranch() string {
	todo("DefaultBaseBranch")
	return "fake-default-base-branch"
}

// PullRequest uses the current branch to create a pull request on GitHub.
func PullRequest(issue issues.Issue) {
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

}

// InitRepo initializes a Git repo.
func InitRepo() error {
	todo("InitRepo")
	return nil
}

// AddAll stages all changes.
func AddAll() error {
	todo("AddAll")
	return nil
}

// Commit the changes.
func Commit(message string) error {
	todo("Commit")
	return nil
}

// Push updates to the remote repo.
func Push(branchName string) error {
	todo("Push")
	return nil
}
