package cmd

import (
	"io/ioutil"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/paultyng/chaki/server"
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
		tasks, err := loadTaskConfig(cmd)
		if err != nil {
			return err
		}

		keyFile, err := os.Open(viper.GetString("jwt-key-file"))
		if err != nil {
			return err
		}
		keyBytes, err := ioutil.ReadAll(keyFile)
		if err != nil {
			return err
		}
		signKey, err := jwt.ParseRSAPrivateKeyFromPEM(keyBytes)
		if err != nil {
			return err
		}

		s := server.New(&server.Config{
			Tasks:             tasks,
			PrivateKey:        signKey,
			OAuthClientID:     viper.GetString("oauth-client-id"),
			OAuthClientSecret: viper.GetString("oauth-client-secret"),
		})

		return s.Start(viper.GetString("service-bind"))
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)

	serveCmd.PersistentFlags().StringP("bind", "b", ":3000", "binding for the server")
	_ = viper.BindPFlag("service-bind", serveCmd.PersistentFlags().Lookup("bind"))
	viper.SetDefault("service-bind", ":3000")

	serveCmd.PersistentFlags().StringP("key", "k", "chaki.key", "jwt private key file")
	_ = viper.BindPFlag("jwt-key-file", serveCmd.PersistentFlags().Lookup("key"))
	viper.SetDefault("jwt-key-file", "chaki.key")

	serveCmd.PersistentFlags().String("client-id", "", "OAuth2 client ID")
	_ = viper.BindPFlag("oauth-client-id", serveCmd.PersistentFlags().Lookup("client-id"))

	serveCmd.PersistentFlags().String("client-secret", "", "OAuth2 client secret")
	_ = viper.BindPFlag("oauth-client-secret", serveCmd.PersistentFlags().Lookup("client-secret"))

	addTaskConfigFlags(serveCmd)
}
