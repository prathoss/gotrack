package cmd

import (
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Command to run action",
}

func init() {
	rootCmd.AddCommand(runCmd)
}
