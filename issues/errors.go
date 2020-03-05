package issues

import "fmt"

// BranchHasNoIDError represents a branch name that does not have the correct format for parsing.
type BranchHasNoIDError struct {
	branchName string
}

// Error displays the error messag as a string.
func (e *BranchHasNoIDError) Error() string {
	return fmt.Sprintf("\"%s\" branch has no issue ID", emphasis(e.branchName))
}
