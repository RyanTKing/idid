package main

import (
	"fmt"
	"os"
)

func die(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
	fmt.Println()
	os.Exit(1)
}

func main() {
	rootCmd.AddCommand(thisCmd)
	if err := rootCmd.Execute(); err != nil {
		die("error: %s", err.Error())
	}
}
