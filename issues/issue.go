package issues

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/fatih/color"
)

// TODO: REMOVE ME
func todo(message string) {
	fmt.Println(color.YellowString("TODO:"), fmt.Sprintf("Implement issues.%s", message))
}

// Issue stores the data for the current issue.
type Issue struct {
	ID    string
	title string
}

// String representation of an issue.
func (i *Issue) String() string {
	return fmt.Sprintf("%s: %s", i.ID, i.Title())
}

// BranchName from issue ID and title.
// Ref: https://github.com/lakshmichandrakala/go-parameterize
func (i Issue) BranchName() string {
	reAlphaNum := regexp.MustCompile("[^A-Za-z0-9]+")
	reTrim := regexp.MustCompile("^-|-$")

	title := reAlphaNum.ReplaceAllString(i.title, "-")
	title = reTrim.ReplaceAllString(title, "")

	id := reAlphaNum.ReplaceAllString(i.ID, "-")
	id = reTrim.ReplaceAllString(id, "")

	branch := strings.Join([]string{id, title}, "-")

	return strings.ToLower(branch)
}

// Title returns the titleized form of the issue title.
// Ref: https://golangcookbook.com/chapters/strings/title/
func (i Issue) Title() string {
	words := strings.Fields(i.title)
	smallwords := " a an on of the to "

	for index, word := range words {
		if strings.Contains(smallwords, " "+word+" ") {
			words[index] = word
		} else {
			words[index] = strings.Title(word)
		}
	}
	return strings.Join(words, " ")
}

// FromBranch gets issue info from the branch name.
func FromBranch(branchName string) (Issue, error) {
	todo("FromBranch")
	return Issue{ID: "FAKE-123", title: "fake yet really long title"}, nil
}

// FromTracker gets issue info by making an HTTP request to the issue tracker API.
func FromTracker(issueID string) (Issue, error) {
	todo("FromTracker")
	return Issue{ID: issueID, title: "fake yet really long title"}, nil
}

// URL returns the URL for the issue.
func URL(issueID string) string {
	todo("URL")
	return fmt.Sprintf("https://example.com/issues/%s", issueID)
}
