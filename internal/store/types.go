package store

import "time"

// Issue is a single GitHub issue (or pull request)
type Issue struct {
	Shorthand string
	URL       string
}

// Entry is a single entry in the data store
type Entry struct {
	Msg     string
	Issues  []Issue
	Created time.Time
}
