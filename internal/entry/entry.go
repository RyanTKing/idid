package entry

import (
	"strings"
	"time"

	"github.com/RyanTKing/idid/internal/git"
)

// Possible types of entries
const (
	EntryNote   EntryType = "NOTE"
	EntryIssue            = "ISSUE"
	EntryPull             = "PULL"
	EntryClosed           = "CLOSED"
	EntryMerged           = "MERGED"

	maxTypeLen = 6
)

// New returns a new entry
func New(created time.Time, msg string, issueShorthands ...string) (*Entry, error) {
	issues := []Issue{}
	entryType := EntryNote
	for _, shorthand := range issueShorthands {
		url, err := git.ExpandLink(shorthand)
		if err != nil {
			return nil, err
		}
		issue := Issue{
			Shorthand: shorthand,
			URL:       url,
		}
		if strings.Contains(url, "issue") {
			entryType = EntryIssue
		}
		if strings.Contains(url, "pull") {
			entryType = EntryPull
		}
		issues = append(issues, issue)
	}
	entry := &Entry{
		Msg:     msg,
		Issues:  issues,
		Created: created,
		Type:    entryType,
	}

	return entry, nil
}