package main

import (
	"fmt"

	"github.com/RyanTKing/idid/internal/store"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	linkColor = color.New(color.FgBlue).Add(color.Underline)

	whatCmd = &cobra.Command{
		Use:   "what",
		Short: "See the what you've done in the past day",
		Long:  `See all your log messages since the previous day`,
		Run: func(cmd *cobra.Command, args []string) {
			entries, err := store.Read(0, 1)
			if err != nil {
				die("error: %s", err.Error())
			}

			for i := len(entries) - 1; i >= 0; i-- {
				entry := entries[i]
				created := entry.Created.Format("Monday, Jan 2 3:04 PM")
				fmt.Printf("%s: %s ", created, entry.Msg)
				for i, issue := range entry.Issues {
					linkColor.Printf("%s", issue.Shorthand)
					if i < len(entry.Issues)-1 {
						fmt.Printf(", ")
					}
				}
				fmt.Println()
			}
		},
	}
)
