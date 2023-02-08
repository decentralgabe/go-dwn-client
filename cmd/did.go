package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(didCmd)
}

var (
	didCmd = &cobra.Command{
		Use:   "did",
		Short: "Interact with the dids",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}
)
