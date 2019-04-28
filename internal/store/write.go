package store

import (
	"encoding/json"
	"os"
	"time"

	"github.com/ryantking/idid/internal/entry"
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

func write(now time.Time, dir, msg string, issues ...string) error {
	path := getPath(now, dir)
	f, err := os.Open(path)
	entries := []*entry.Entry{}
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

	entry, err := entry.New(now, msg, issues...)
	if err != nil {
		return err
	}
	entries = append(entries, entry)

	f, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer f.Close()
	return json.NewEncoder(f).Encode(entries)
}

func WriteClose(issue string) error {
	now := time.Now()
	dir, err := getDirectory()
	if err != nil {
		return err
	}

	return writeClose(now, dir, issue)
}

func writeClose(now time.Time, dir, issue string) error {
	path := getPath(now, dir)
	f, err := os.Open(path)
	entries := []*entry.Entry{}
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

	entry, err := entry.NewClose(now, issue)
	if err != nil {
		return err
	}
	entries = append(entries, entry)

	f, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer f.Close()
	return json.NewEncoder(f).Encode(entries)
}

func WriteMerge(issue string) error {
	now := time.Now()
	dir, err := getDirectory()
	if err != nil {
		return err
	}

	return writeMerge(now, dir, issue)
}

func writeMerge(now time.Time, dir, issue string) error {
	path := getPath(now, dir)
	f, err := os.Open(path)
	entries := []*entry.Entry{}
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

	entry, err := entry.NewMerge(now, issue)
	if err != nil {
		return err
	}
	entries = append(entries, entry)

	f, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer f.Close()
	return json.NewEncoder(f).Encode(entries)
}
