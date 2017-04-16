package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.ua-ecm.com/chaki/server"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:        "serve",
	Aliases:    []string{"s"},
	SuggestFor: []string{"web"},
	Short:      "Starts the web server",
	Long: `Serve starts the web server for both static UI
files and the tasks API endpoints at the specified binding.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := loadTaskConfig(cmd)
		if err != nil {
			return err
		}

		s := server.New(config)

		return s.Start(viper.GetString("service-bind"))
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)

	serveCmd.PersistentFlags().StringP("bind", "b", ":3000", "binding for the server")
	viper.BindPFlag("service-bind", serveCmd.PersistentFlags().Lookup("bind"))
	viper.SetDefault("service-bind", ":3000")

	addTaskConfigFlags(serveCmd)
}
