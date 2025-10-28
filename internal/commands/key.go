package commands

import (
	"fmt"
	"os"

	"github.com/gigvault/shared/pkg/crypto"
	"github.com/spf13/cobra"
)

// NewKeyCommand returns the key management command
func NewKeyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "key",
		Short: "Key management operations",
		Long:  `Generate and manage cryptographic keys`,
	}

	cmd.AddCommand(&cobra.Command{
		Use:   "generate [name]",
		Short: "Generate a new ECDSA key pair",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			outputDir, _ := cmd.Flags().GetString("output")
			curve, _ := cmd.Flags().GetString("curve")

			fmt.Printf("Generating %s key for %s...\n", curve, name)

			var keyPEM []byte
			var err error

			switch curve {
			case "P-256":
				key, err := crypto.GenerateP256Key()
				if err != nil {
					return fmt.Errorf("failed to generate key: %w", err)
				}
				keyPEM, err = crypto.EncodePrivateKeyToPEM(key)
			case "P-384":
				key, err := crypto.GenerateP384Key()
				if err != nil {
					return fmt.Errorf("failed to generate key: %w", err)
				}
				keyPEM, err = crypto.EncodePrivateKeyToPEM(key)
			default:
				return fmt.Errorf("unsupported curve: %s", curve)
			}

			// Save private key
			if err != nil {
				return err
			}

			keyPath := fmt.Sprintf("%s/%s.key", outputDir, name)
			if err := os.MkdirAll(outputDir, 0755); err != nil {
				return err
			}
			if err := os.WriteFile(keyPath, keyPEM, 0600); err != nil {
				return err
			}

			fmt.Printf("Private key saved to: %s\n", keyPath)

			return nil
		},
	})

	cmd.PersistentFlags().StringP("output", "o", ".", "Output directory")
	cmd.PersistentFlags().StringP("curve", "c", "P-256", "Elliptic curve (P-256, P-384)")

	return cmd
}
