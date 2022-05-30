/*
Package cmd
Copyright © 2022 Mehdi Teymorian

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func Execute() {
	rootCmd := &cobra.Command{
		Use:   "qsse-test",
		Short: "stress test for QSSE",
		Long:  `a service for stress testing QSSE including generator, server, and subscriber.`,
	}

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
