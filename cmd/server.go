package cmd

import (
	"github.com/spf13/cobra"
	"gotrack/pkg"
	"gotrack/server"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start server",
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := pkg.NewAppConfigFromViper()
		if err != nil {
			return err
		}

		dc, err := pkg.NewDependencyContainer(config)
		if err != nil {
			return err
		}

		return server.Run(dc)
	},
}

func init() {
	runCmd.AddCommand(serverCmd)
}
