package issues

import (
	"github.com/fatih/color"
)

// Issue stores the data for the current issue.
type Issue struct {
	ID string
}

// BranchName from issue ID and title.
func (i Issue) BranchName() string {
	color.Yellow("TODO: Implement BranchName") // REMOVE ME
	return "fake-branch-name"
}

// FromTracker gets issue info by making an HTTP request to the issue tracker API.
func FromTracker(issueID string) (*Issue, error) {
	color.Yellow("TODO: Implement FromTracker") // REMOVE ME
	return &Issue{ID: issueID}, nil
}
