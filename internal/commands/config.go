package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewConfigCommand returns the configuration management command
func NewConfigCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Configuration operations",
		Long:  `Manage GigVault CLI configuration`,
	}

	cmd.AddCommand(&cobra.Command{
		Use:   "show",
		Short: "Show current configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("GigVault Configuration:")
			fmt.Println("  CA URL: http://localhost:8080")
			fmt.Println("  RA URL: http://localhost:8081")
			fmt.Println("  Auth: Not configured")
			return nil
		},
	})

	return cmd
}
