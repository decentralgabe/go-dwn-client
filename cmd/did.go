package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(didCmd)
	rootCmd.PersistentFlags().StringVar(&didDID, "did", "", "did to use for the command")
}

var (
	didDID string

	didCmd = &cobra.Command{
		Use:   "did",
		Short: "Interact with the dids",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}
)
