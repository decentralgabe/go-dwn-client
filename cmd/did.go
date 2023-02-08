package cmd

import (
	"errors"

	ssicrypto "github.com/TBD54566975/ssi-sdk/crypto"
	"github.com/spf13/cobra"

	"github.com/decentralgabe/go-rdr-client/internal"
)

func init() {
	rootCmd.AddCommand(didCmd)
	rootCmd.PersistentFlags().StringVar(&didDID, "did", "", "did to use for the command")

	didCmd.AddCommand(didViewCmd)
	didCmd.AddCommand(didGenerateCmd)
	didCmd.AddCommand(didAddCmd)
	didCmd.AddCommand(didRemoveCmd)
}

var (
	didDID string

	didTable = new(internal.DIDTable)

	didCmd = &cobra.Command{
		Use:   "did",
		Short: "Interact with the dids",
		Run: func(cmd *cobra.Command, args []string) {
			didTable.PrintDIDs()
		},
	}

	didViewCmd = &cobra.Command{
		Use:   "view",
		Short: "Interact with the did table",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				return cmd.Help()
			}
			didTable.PrintDID(didDID)
			return nil
		},
	}

	didGenerateCmd = &cobra.Command{
		Use:   "generate",
		Short: "Generate a new did",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				return cmd.Help()
			}
			keyPair, err := internal.GenerateDIDKeyPair(ssicrypto.Ed25519)
			if err != nil {
				return err
			}
			if err := didTable.AddDID(*keyPair); err != nil {
				return err
			}
			return nil
		},
	}

	didAddCmd = &cobra.Command{
		Use:   "add",
		Short: "Add a did to the did table",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 3 {
				return cmd.Help()
			}
			if didDID == "" {
				return errors.New("did is a required flag")
			}
			keyPair := internal.KeyPair{
				ID:              didDID,
				KeyType:         ssicrypto.KeyType(args[0]),
				PublicKeyBase58: args[1],
			}
			if args[2] != "" {
				keyPair.PrivateKeyBase58 = args[2]
			}
			return didTable.AddDID(keyPair)
		},
	}

	didRemoveCmd = &cobra.Command{
		Use:   "remove",
		Short: "Remove a did from the did table",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				return cmd.Help()
			}
			if didDID == "" {
				return errors.New("did is a required flag")
			}
			return didTable.RemoveDID(didDID)
		},
	}
)
