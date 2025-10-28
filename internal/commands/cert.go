package commands

import (
	"crypto/x509/pkix"
	"fmt"
	"os"

	"github.com/gigvault/shared/pkg/crypto"
	"github.com/spf13/cobra"
)

// NewCertCommand returns the certificate management command
func NewCertCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cert",
		Short: "Certificate operations",
		Long:  `Manage certificates: issue, revoke, list, inspect`,
	}

	cmd.AddCommand(&cobra.Command{
		Use:   "create [common-name]",
		Short: "Create a self-signed certificate",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cn := args[0]
			outputDir, _ := cmd.Flags().GetString("output")
			validityDays, _ := cmd.Flags().GetInt("validity")

			fmt.Printf("Creating certificate for %s...\n", cn)

			// Generate key
			key, err := crypto.GenerateP256Key()
			if err != nil {
				return fmt.Errorf("failed to generate key: %w", err)
			}

			// Create certificate template
			template, err := crypto.CreateCertificateTemplate(cn, validityDays)
			if err != nil {
				return fmt.Errorf("failed to create template: %w", err)
			}

			template.Subject = pkix.Name{CommonName: cn}

			// Self-sign
			certPEM, err := crypto.SignCertificate(template, template, &key.PublicKey, key)
			if err != nil {
				return fmt.Errorf("failed to sign certificate: %w", err)
			}

			// Save certificate
			certPath := fmt.Sprintf("%s/%s.crt", outputDir, cn)
			if err := os.MkdirAll(outputDir, 0755); err != nil {
				return err
			}
			if err := os.WriteFile(certPath, certPEM, 0644); err != nil {
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

			fmt.Printf("Certificate saved to: %s\n", certPath)
			fmt.Printf("Private key saved to: %s\n", keyPath)

			return nil
		},
	})

	cmd.PersistentFlags().StringP("output", "o", ".", "Output directory")
	cmd.PersistentFlags().IntP("validity", "v", 365, "Validity in days")

	return cmd
}
