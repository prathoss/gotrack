package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"gotrack/pkg"
)

// serverCfgCmd represents the serverCgf command
var serverCfgCmd = &cobra.Command{
	Use:   "serverCfg",
	Short: "Generate empty server configuration file template",
	RunE: func(cmd *cobra.Command, args []string) error {
		tmplt, err := pkg.GenerateAppConfig()
		if err != nil {
			return err
		}
		fmt.Println(tmplt)
		return nil
	},
}

func init() {
	generateCmd.AddCommand(serverCfgCmd)
}
