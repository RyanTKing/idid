package store

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"os"

	"github.com/RyanTKing/idid/internal/config"
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
	entries := []Entry{}
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

	issues := []Issue{}
	for _, shorthand := range issueShorthands {
		url, err := git.ExpandLink(shorthand)
		if err != nil {
			return err
		}
		issue := Issue{
			Shorthand: shorthand,
			URL:       url,
		}
		issues = append(issues, issue)
	}
	entry := Entry{
		Msg:     msg,
		Issues:  issues,
		Created: now,
	}
	entries = append(entries, entry)

	f, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer f.Close()
	return json.NewEncoder(f).Encode(entries)
}

// Read reads the messages written in the last 24 hours and returns the entry
func Read(sinceDays, sinceMonths int) ([]Entry, error) {
	now := time.Now()
	since := now.AddDate(0, -1*sinceMonths, -1*sinceDays)
	dir, err := getDirectory()
	if err != nil {
		return []Entry{}, err
	}

	return read(now, since, dir)
}

func read(now, since time.Time, dir string) ([]Entry, error) {
	entries := []Entry{}
	for t := now; !t.Before(since); t = t.AddDate(0, 0, -1) {
		path := getPath(t, dir)
		f, err := os.Open(path)
		defer f.Close()
		if os.IsNotExist(err) {
			continue
		} else if err != nil {
			return []Entry{}, err
		}

		newEntries := []Entry{}
		err = json.NewDecoder(f).Decode(&newEntries)
		if err != nil {
			return []Entry{}, err
		}
		entries = append(entries, newEntries...)
	}

	return entries, nil
}

func getPath(now time.Time, dir string) string {
	y, m, d := now.Date()
	return fmt.Sprintf("%s/%2d%02d%02d.json", dir, y, m, d)
}

func getDirectory() (string, error) {
	cfg := config.Get()
	subDir := strings.Split(cfg.GitHub.Endpoint, "//")[1]
	subDir = strings.Replace(subDir, ".", "_", -1)
	dir := fmt.Sprintf("%s/%s", cfg.StorageDir, subDir)
	if err := checkDirectory(dir); err != nil {
		return "", err
	}

	return dir, nil
}

func checkDirectory(dir string) error {
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return os.MkdirAll(dir, os.ModePerm)
	} else if err != nil {
		return err
	}

	return nil
}
