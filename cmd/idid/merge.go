package main

import (
	"github.com/ryantking/idid/internal/store"
	"github.com/spf13/cobra"
)

var (
	mergeCmd = &cobra.Command{
		Use:   "merge",
		Short: "Record that you merged a pull",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				die("please provide an pull")
			}

			issue := args[0]
			err := store.WriteMerge(issue)
			if err != nil {
				die("error: %s", err.Error())
			}
		},
	}
)
