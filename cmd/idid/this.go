package main

import (
	"github.com/ryantking/idid/internal/store"
	"github.com/spf13/cobra"
)

var (
	thisCmd = &cobra.Command{
		Use:   "this",
		Short: "Record something you did",
		Long:  `Record something you did linking it to any corresponding issues or pull requests`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				die("please provide a message")
			}

			msg := args[0]
			issues := args[1:len(args)]
			err := store.Write(msg, issues...)
			if err != nil {
				die("error: %s", err.Error())
			}
		},
	}
)
