package cmd

import (
	"github.com/spf13/cobra"

	"github.com/decentralgabe/go-rdr-client/pkg"
)

var (
	routeTable *pkg.RouteTable

	routeCmd = &cobra.Command{
		Use:   "route",
		Short: "Interact with the route table",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
)
