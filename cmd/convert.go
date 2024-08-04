package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(convertCmd)
}

var convertCmd = &cobra.Command{
	Use: "convert [filename]",
	Long: `Takes in a Mongo query through a file or command line
and converts it into Go code to be used for a go mongo-driver.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fileName := args[0]
		fileBytes, err := os.ReadFile(fileName)
		if err != nil {
			return err
		}
		os.Stdout.Write(fileBytes)
		return nil
	},
}
