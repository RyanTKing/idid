package main

import "log"

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("error: %s", err.Error())
	}
}
