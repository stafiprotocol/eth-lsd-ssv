package cmd

import (
	"context"
	"os"

	"github.com/spf13/cobra"
)

var (
	appName = "lsd-ssv-client"
)

const (
	flagKeystorePath = "keystore_path"
	flagLogLevel     = "log_level"
	flagConfigPath   = "config"
	flagViewMode     = "view_mode"

	defaultKeystorePath = "./keys"
	defaultConfigPath   = "./config.toml"
)

// NewRootCmd returns the root command.
func NewRootCmd() *cobra.Command {
	// RootCmd represents the base command when called without any subcommands
	var rootCmd = &cobra.Command{
		Use:   appName,
		Short: "lsd-ssv-client",
	}

	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, _ []string) error {

		return nil
	}

	rootCmd.AddCommand(
		importAccountCmd(),
		importMnemonicCmd(),
		startSsvCmd(),
		versionCmd(),
	)
	return rootCmd
}

func Execute() {
	cobra.EnableCommandSorting = false

	rootCmd := NewRootCmd()
	rootCmd.SilenceUsage = true
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	ctx := context.Background()

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}
