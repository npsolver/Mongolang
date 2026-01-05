package cmd

import (
	"fmt"
	"os"

	"github.com/npsolver/Mongolang/codegen"
	"github.com/npsolver/Mongolang/parser"
	"github.com/npsolver/Mongolang/scanner"
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

		// Input
		fileName := args[0]
		fileBytes, err := os.ReadFile(fileName)
		if err != nil {
			return err
		}

		fmt.Println("\n\nPrinting Scanned symbols\n\n")

		// Scanning
		symbols, err := scanner.Scan(string(fileBytes))
		if err != nil {
			return err
		}
		for _, tk := range symbols {
			tk.Print()
		}

		fmt.Println("\n\nPrinting Parsed data\n\n")

		// Parsing
		startNode := parser.Parse(symbols)

		// Context Sensitive Analysis
		// Stuff like type checking
		// Todo
		// also test stuff in parseCFS 
		// like rn we are transitioning until the end
		// make sure to find whether we can come back

		// Code Generation
		filter := codegen.Generate(startNode)

		fmt.Println(filter)

		return nil
	},
}
