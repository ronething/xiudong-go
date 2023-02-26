package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	version = "unknown"
	commit  = "?"
	date    = "?"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "check version",
	Long:  `check showstart cli version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(os.Stdout, `git commit hash: %s
build timestamp: %s
version: %s
`, commit, date, version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
