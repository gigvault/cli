package commands

import (
	"crypto/x509/pkix"
	"fmt"
	"os"

	"github.com/gigvault/shared/pkg/crypto"
	"github.com/spf13/cobra"
)

// NewCSRCommand returns the CSR management command
func NewCSRCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "csr",
		Short: "Certificate Signing Request operations",
		Long:  `Create and manage Certificate Signing Requests`,
	}

	cmd.AddCommand(&cobra.Command{
		Use:   "create [common-name]",
		Short: "Create a new CSR",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cn := args[0]
			outputDir, _ := cmd.Flags().GetString("output")

			fmt.Printf("Creating CSR for %s...\n", cn)

			// Generate key
			key, err := crypto.GenerateP256Key()
			if err != nil {
				return fmt.Errorf("failed to generate key: %w", err)
			}

			// Create CSR
			subject := pkix.Name{
				CommonName: cn,
			}

			csrPEM, err := crypto.CreateCSR(key, subject)
			if err != nil {
				return fmt.Errorf("failed to create CSR: %w", err)
			}

			// Save CSR
			csrPath := fmt.Sprintf("%s/%s.csr", outputDir, cn)
			if err := os.MkdirAll(outputDir, 0755); err != nil {
				return err
			}
			if err := os.WriteFile(csrPath, csrPEM, 0644); err != nil {
				return err
			}

			// Save key
			keyPEM, err := crypto.EncodePrivateKeyToPEM(key)
			if err != nil {
				return err
			}
			keyPath := fmt.Sprintf("%s/%s.key", outputDir, cn)
			if err := os.WriteFile(keyPath, keyPEM, 0600); err != nil {
				return err
			}

			fmt.Printf("CSR saved to: %s\n", csrPath)
			fmt.Printf("Private key saved to: %s\n", keyPath)

			return nil
		},
	})

	cmd.PersistentFlags().StringP("output", "o", ".", "Output directory")

	return cmd
}
