package git

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/greganswer/mgit_go/issues"
)

// CurrentBranch returns the current branch for this Git repo.
func CurrentBranch() string {
	color.Yellow("TODO: Implement git.CurrentBranch") // REMOVE ME
	return "fake-current-branch"
}

// DefaultBaseBranch returns the base branch for this Git repo.
func DefaultBaseBranch() string {
	color.Yellow("TODO: Implement git.DefaultBaseBranch") // REMOVE ME
	return "fake-default-base-branch"
}

// PullRequest uses the current branch to create a pull request on GitHub.
func PullRequest(issue issues.Issue) {
	color.Yellow("TODO: Implement git.PullRequest") // REMOVE ME
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

	body := fmt.Sprintf(template, issue.Title(), issueTracker, issue.ID, issue.URL())

	fmt.Println(body) // TODO: REMOVE ME
}
