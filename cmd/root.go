package cmd

import (
	"github.com/mehditeymorian/qsse-test/internal/cmd/generate"
	"github.com/mehditeymorian/qsse-test/internal/cmd/serve"
	"github.com/spf13/cobra"
)

// Execute is the entry point for the application.
func Execute() {
	rootCmd := &cobra.Command{ //nolint:exhaustruct
		Use:   "qsse-test",
		Short: "stress test for QSSE",
		Long:  `a service for stress testing QSSE including generator, server, and subscriber.`,
	}

	rootCmd.AddCommand(
		serve.Command(),
		generate.Command(),
	)

	panic(rootCmd.Execute())
}
