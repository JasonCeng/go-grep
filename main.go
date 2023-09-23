package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-grep/grep"
	"go-grep/version"
	"os"
)

// TODO
//1、Write functionality to identify patterns in a line of text
//2、Implement a way to handle both input files and shell input
//3、Create a function to display matched lines
//4、Process command line flags and arguments
//5、Use Goroutines to concurrently process data

var (
	pattern       string
	fileNames     []string
	concurrentNum int
)

// The main command describes the service and
// defaults to printing the help message.
var mainCmd = &cobra.Command{
	Use:   "go-grep",
	Short: "Search for a pattern in files.",
	RunE: func(cmd *cobra.Command, args []string) error {
		for _, file := range fileNames {
			err := grep.Grep(pattern, file, concurrentNum)
			if err != nil {
				fmt.Errorf("Error searching %s:%v\n", file, err)
			}
			return err
		}
		return nil
	},
}

// TODO: 支持管道输入
func main() {
	//// Define command-line flags that are valid for all go-grep commands and
	//// subcommands.
	//mainFlags := mainCmd.PersistentFlags()

	mainCmd.AddCommand(version.Cmd())

	mainCmd.Flags().StringVarP(&pattern, "pattern", "p", "", "Pattern to search for")
	mainCmd.Flags().StringSliceVarP(&fileNames, "file", "f", []string{}, "Files to search in (can specify multiple)")
	mainCmd.Flags().IntVarP(&concurrentNum, "concurrent", "c", 1, "Number of concurrent goroutines")

	// On failure Cobra prints the usage message and error string, so we only
	// need to exit with a non-0 status
	if err := mainCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
