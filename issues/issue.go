package issues

import (
	"fmt"
	"regexp"
	"strconv"
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
	if i.ID != "" && i.Title() != "" {
		return fmt.Sprintf("%s: %s", i.ID, i.Title())
	}
	return i.Title()
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

// Title is the titleized form of the issue title. Uses O(n) time and space.
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

// APIURL is the URL to retrieve issue data via the issue tracker API.
func (i Issue) APIURL() string {
	todo("APIURL")
	return fmt.Sprintf("https://api.example.com/issues/%s", i.ID)
}

// WebURL is the URL to retrieve issue data via the issue tracker website.
func (i Issue) WebURL() string {
	todo("WebURL")
	return fmt.Sprintf("https://example.com/issues/%s", i.ID)
}

// APIURL is the URL to retrieve issue data via the issue tracker API.
func APIURL(issueID string) string {
	todo("APIURL")
	return fmt.Sprintf("https://api.example.com/issues/%s", issueID)
}

// FromBranch creates an issue from the branch name. Uses O(n) time and space.
func FromBranch(branchName string) (Issue, error) {
	parts := strings.Split(branchName, "-")
	for i, part := range parts {
		if _, err := strconv.Atoi(part); err == nil {
			issue := Issue{
				ID:    strings.Join(parts[:i+1], "-"),
				title: strings.Join(parts[i+1:], "-"),
			}
			return issue, nil
		}
	}
	return Issue{}, &BranchHasNoIDError{branchName}
}

// FromTracker creates an issue by making an HTTP request to the issue tracker API.
func FromTracker(issueID string) (Issue, error) {
	todo("FromTracker")
	return Issue{ID: issueID, title: "fake yet really long title"}, nil
}

func emphasis(message string) string {
	return color.CyanString(message)
}
