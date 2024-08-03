package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "leapher",
		Short: "A tool to generate Golang code from Mongo Queries.",
		Long: `leapher is a command line tool that can be used to convert queries
written for MongoDB to Golang code to be used with the mongo-driver. 
It handles the tedious work of converting the queries for you.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
