package entry

import (
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
)

// Print prints out the entry to standard out
func (e *Entry) Print() {
	e.Fprint(os.Stdout)
}

// Fprint prints the entry to the given io.Writer
func (e *Entry) Fprint(w io.Writer) {
	e.printType(w)
	e.printTS(w)
	e.printMsg(w)
	e.printIssues(w)
	fmt.Fprintf(w, "\n")
}

func (e *Entry) printIssues(w io.Writer) {
	c := color.New(color.FgCyan).Add(color.Underline)
	for i, issue := range e.Issues {
		c.Fprintf(w, "%s", issue.Shorthand)
		if i < len(e.Issues)-1 {
			fmt.Fprintf(w, ", ")
		}
	}
}

func (e *Entry) printMsg(w io.Writer) {
	fmt.Fprintf(w, "%s ", e.Msg)
}

func (e *Entry) printType(w io.Writer) {
	var c *color.Color
	switch e.Type {
	case EntryIssue:
		c = color.New(color.FgBlue)
	case EntryPull:
		c = color.New(color.FgGreen)
	case EntryClosed:
		c = color.New(color.FgRed)
	case EntryMerged:
		c = color.New(color.FgMagenta)
	default:
		c = color.New(color.FgWhite)
	}

	c.Fprint(w, "[")
	c.Fprint(w, e.Type)
	c.Fprint(w, "]")
	fmt.Fprint(w, " ")
}

func (e *Entry) printTS(w io.Writer) {
	ts := e.Created.Format("Mon, Jan 2 3:04 PM")
	fmt.Fprintf(w, "%s | ", ts)

}
