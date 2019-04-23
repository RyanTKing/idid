package main

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "idid",
	Short: "idid is a fast and simple way to log  your work",
	Long:  `idid is a simple tool that allows you to log what you're spending and link it to Githb issues`,
}
