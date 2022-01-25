package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gotrack/pkg"
	"gotrack/pkg/applogger"
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

		dev := viper.GetBool("dev")
		applogger.Init(dev)

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
