package main

import (
	"github.com/RyanTKing/idid/internal/store"
	"github.com/spf13/cobra"
)

var (
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
				entries[i].Print()
			}
		},
	}
)
