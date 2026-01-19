package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/nerveband/beeper-cli/internal/config"
	"github.com/nerveband/beeper-cli/internal/output"
)

var (
	cfg          *config.Config
	outputFormat string
)

var rootCmd = &cobra.Command{
	Use:   "beeper",
	Short: "Beeper CLI - Command-line interface for Beeper Desktop API",
	Long: `A cross-platform CLI for the Beeper Desktop API.
Provides LLM-friendly interfaces for reading and sending messages
across all Beeper-supported chat networks.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		var err error
		cfg, err = config.Load()
		if err != nil {
			return fmt.Errorf("failed to load config: %w", err)
		}

		// Override output format if flag is set
		if outputFormat != "" {
			cfg.OutputFormat = outputFormat
		}

		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&outputFormat, "output", "o", "", "Output format (json, text, markdown)")
}

// getOutputFormat returns the configured output format
func getOutputFormat() output.Format {
	return output.Format(cfg.OutputFormat)
}
