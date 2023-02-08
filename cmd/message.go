package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(msgCmd)
}

var (
	msgCmd = &cobra.Command{
		Use:   "msg",
		Short: "Send messages with dids",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
)
