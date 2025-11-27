package cmd

import (
	// "github.com/Npsolver/Mongolang/internal/compiler/scanner"

	"github.com/npsolver/Mongolang/internal/lib/dfa"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(convertCmd)
}

var convertCmd = &cobra.Command{
	Use:  "convert [filename]",
	Long: `Takes in a Mongo query through a file or command line and converts it into Go code to be used for a go mongo-driver.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fileName := args[0]

		// Read the query from file
		text := <-dfa.NewFileReader(fileName)
		print(text)

		// Scan into tokens
		

		// Parse into a tree

		// Context-Sensitive Analysis

		// Code generation

		return nil
	},
}
