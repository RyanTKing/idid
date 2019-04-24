package store

import (
	"encoding/json"
	"os"
	"time"

	"github.com/RyanTKing/idid/internal/entry"
)

// Read reads the messages written in the last 24 hours and returns the entry
func Read(sinceDays, sinceMonths int) ([]entry.Entry, error) {
	now := time.Now()
	since := now.AddDate(0, -1*sinceMonths, -1*sinceDays)
	dir, err := getDirectory()
	if err != nil {
		return []entry.Entry{}, err
	}

	return read(now, since, dir)
}

func read(now, since time.Time, dir string) ([]entry.Entry, error) {
	entries := []entry.Entry{}
	for t := now; !t.Before(since); t = t.AddDate(0, 0, -1) {
		path := getPath(t, dir)
		f, err := os.Open(path)
		defer f.Close()
		if os.IsNotExist(err) {
			continue
		} else if err != nil {
			return []entry.Entry{}, err
		}

		newEntries := []entry.Entry{}
		err = json.NewDecoder(f).Decode(&newEntries)
		if err != nil {
			return []entry.Entry{}, err
		}
		entries = append(entries, newEntries...)
	}

	return entries, nil
}
