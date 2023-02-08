package cmd

import (
	"errors"

	"github.com/spf13/cobra"

	"github.com/decentralgabe/go-rdr-client/internal"
)

func init() {
	rootCmd.AddCommand(routeCmd)
	rootCmd.PersistentFlags().StringVar(&did, "did", "", "did to use for the command")

	routeCmd.AddCommand(routeViewCmd)
	routeCmd.AddCommand(routeAddCmd)
	routeCmd.AddCommand(routeRemoveCmd)
}

var (
	did string

	routeTable = new(internal.RouteTable)

	routeCmd = &cobra.Command{
		Use:   "route",
		Short: "Interact with the route table",
		Run: func(cmd *cobra.Command, args []string) {
			routeTable.PrintRoutes()
		},
	}

	routeViewCmd = &cobra.Command{
		Use:   "view",
		Short: "View the route table for a given DID",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				return cmd.Help()
			}
			routeTable.PrintRoute(did)
			return nil
		},
	}

	routeAddCmd = &cobra.Command{
		Use:   "add",
		Short: "Add a route to the route table",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return cmd.Help()
			}
			if did == "" {
				return errors.New("did is a required flag")
			}
			return routeTable.AddRoute(did, args[0])
		},
	}

	routeRemoveCmd = &cobra.Command{
		Use:   "remove",
		Short: "Remove a route from the route table",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				return cmd.Help()
			}
			if did == "" {
				return errors.New("did is a required flag")
			}
			return routeTable.RemoveRoute(did)
		},
	}
)
