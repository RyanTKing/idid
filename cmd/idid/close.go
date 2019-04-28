package main

import (
	"github.com/ryantking/idid/internal/store"
	"github.com/spf13/cobra"
)

var (
	closeCmd = &cobra.Command{
		Use:   "close",
		Short: "Record that you closed an issue",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				die("please provide an issue")
			}

			issue := args[0]
			err := store.WriteClose(issue)
			if err != nil {
				die("error: %s", err.Error())
			}
		},
	}
)
