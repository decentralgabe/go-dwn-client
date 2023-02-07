package cmd

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/decentralgabe/go-rdr-client/pkg"
)

var (
	// Used for flags.
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "rdr",
		Short: "A tool for interacting with a remote dwn relay",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is config/cli.yaml)")
	rootCmd.AddCommand(routeCmd)
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find current directory.
		current, err := os.Getwd()
		cobra.CheckErr(err)

		// Search config in current directory with name "cli" (without extension).
		viper.AddConfigPath(current)
		viper.AddConfigPath(current + "/config")
		viper.SetConfigType("yaml")
		viper.SetConfigName("cli")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		logrus.Infof("Using config file: %s", viper.ConfigFileUsed())
		routes := viper.Get("routes").(map[string]interface{})
		routeTable = pkg.NewRouteTableFromConfig(routes)
		// routeTable.PrintRoutes()
	}
}
