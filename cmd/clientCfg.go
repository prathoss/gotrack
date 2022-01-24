package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"gotrack/pkg"
)

// clientCfgCmd represents the clientCgf command
var clientCfgCmd = &cobra.Command{
	Use:   "clientCfg",
	Short: "Generate empty client configuration file template",
	RunE: func(cmd *cobra.Command, args []string) error {
		tmplt, err := pkg.GenerateClientConfig()
		if err != nil {
			return err
		}
		fmt.Println(tmplt)
		return nil
	},
}

func init() {
	generateCmd.AddCommand(clientCfgCmd)
}
