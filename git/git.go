package git

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/greganswer/mgit_go/issues"
)

// TODO: REMOVE ME
func todo(message string) {
	fmt.Println(color.YellowString("TODO:"), fmt.Sprintf("Implement git.%s", message))
}

// CurrentBranch returns the current branch for this Git repo.
func CurrentBranch() string {
	todo("CurrentBranch")
	return "fake-current-branch"
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

	body := fmt.Sprintf(template, issue.Title(), issueTracker, issue.ID, issues.URL(issue.ID))

	fmt.Println(body) // TODO: REMOVE ME
}
