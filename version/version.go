package version

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-grep/metadata"
	"runtime"
)

// Program name
const ProgramName = "go-grep"

// Cmd returns the Cobra Command for Version
func Cmd() *cobra.Command {
	return cobraCommand
}

var cobraCommand = &cobra.Command{
	Use:   "version",
	Short: "Print go-grep version.",
	Long:  `Print current version of the go-grep.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			return fmt.Errorf("trailing args detected")
		}
		// Parsing of the command line is done so silence cmd usage
		cmd.SilenceUsage = true
		fmt.Print(GetInfo())
		return nil
	},
}

// GetInfo returns version information for the go-grep
func GetInfo() string {

	return fmt.Sprintf("%s:\n Version: %s\n Commit SHA: %s\n Go version: %s\n"+
		" OS/Arch: %s\n",
		ProgramName, metadata.Version, metadata.CommitSHA, runtime.Version(),
		fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH))
}
