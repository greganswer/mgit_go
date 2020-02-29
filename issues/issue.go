package issues

import (
	"github.com/fatih/color"
)

type Issue struct {
	ID string
}

func (i Issue) BranchName() string {
	color.Yellow("TODO: Implement BranchName") // REMOVE ME
	return "fake-branch-name"
}

func FromTracker(issueID string) (Issue, error) {
	color.Yellow("TODO: Implement FromTracker") // REMOVE ME
	return Issue{ID: issueID}, nil
}
