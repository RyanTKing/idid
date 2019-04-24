package store

import (
	"encoding/json"
	"os"
	"time"

	"github.com/RyanTKing/idid/internal/entry"
	"github.com/RyanTKing/idid/internal/git"
)

// Write writes a message with related issues to the store
func Write(msg string, issues ...string) error {
	now := time.Now()
	dir, err := getDirectory()
	if err != nil {
		return err
	}

	return write(now, dir, msg, issues...)
}

func write(now time.Time, dir, msg string, issueShorthands ...string) error {
	path := getPath(now, dir)
	f, err := os.Open(path)
	entries := []entry.Entry{}
	if err == nil {
		err := json.NewDecoder(f).Decode(&entries)
		if err != nil {
			f.Close()
			return err
		}
		err = os.Remove(path)
		if err != nil {
			f.Close()
			return err
		}
		f.Close()
	} else if !os.IsNotExist(err) {
		return err
	}

	issues := []entry.Issue{}
	for _, shorthand := range issueShorthands {
		url, err := git.ExpandLink(shorthand)
		if err != nil {
			return err
		}
		issue := entry.Issue{
			Shorthand: shorthand,
			URL:       url,
		}
		issues = append(issues, issue)
	}
	entry := entry.Entry{
		Msg:     msg,
		Issues:  issues,
		Created: now,
	}
	entries = append(entries, entry)

	f, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer f.Close()
	return json.NewEncoder(f).Encode(entries)
}
