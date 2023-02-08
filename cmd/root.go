package cmd

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/decentralgabe/go-rdr-client/internal"
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
	// config
	cobra.OnInitialize(initConfig)

	// flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ~/.rdr-cli.json)")
}

func initConfig() {
	// Find the home directory.
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory with name "cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("json")
		viper.SetConfigName(".rdr-cli")
	}

	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err == nil {
		logrus.Infof("Using config file: %s", viper.ConfigFileUsed())
		routeTable = internal.NewRouteTableFromConfig(viper.Get("routes"))
		if logrus.GetLevel() > logrus.InfoLevel {
			routeTable.PrintRoutes()
		}
		didTable = internal.NewDIDTableFromConfig(viper.Get("dids"))
		if logrus.GetLevel() > logrus.InfoLevel {
			didTable.PrintDIDs()
		}
	} else {
		logrus.Warnf("Could not read config file: %s", err.Error())
		configLocation := home + "/.rdr-cli.json"
		if err = os.WriteFile(configLocation, []byte("{}"), 0644); err != nil {
			logrus.Warnf("Could not create config file: %s", err.Error())
		}
		routeTable = internal.NewRouteTableFromConfig(nil)
		didTable = internal.NewDIDTableFromConfig(nil)
	}
}
