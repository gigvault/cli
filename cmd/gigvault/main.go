package main

import (
	"fmt"
	"os"

	"github.com/gigvault/cli/internal/commands"
	"github.com/spf13/cobra"
)

var version = "1.0.0"

func main() {
	rootCmd := &cobra.Command{
		Use:     "gigvault",
		Short:   "GigVault PKI CLI",
		Long:    `Command-line interface for managing GigVault Certificate Authority`,
		Version: version,
	}

	// Add subcommands
	rootCmd.AddCommand(commands.NewCertCommand())
	rootCmd.AddCommand(commands.NewCSRCommand())
	rootCmd.AddCommand(commands.NewKeyCommand())
	rootCmd.AddCommand(commands.NewConfigCommand())

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
