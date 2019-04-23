package store

import (
	"fmt"
	"strings"
	"time"

	"os"

	"github.com/RyanTKing/idid/internal/config"
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
	return nil
}

// Read reads the messages written in the last 24 hours and returns the entry
func Read() ([]Entry, error) {
	now := time.Now()
	dir, err := getDirectory()
	if err != nil {
		return []Entry{}, err
	}

	return read(now, dir)
}

func read(now time.Time, dir string) ([]Entry, error) {
	return []Entry{}, nil
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
