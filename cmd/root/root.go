package root

import (
	"lr-cli/cmd/login"

	"github.com/spf13/cobra"
)

var cfgFile string

func NewRootCmd() *cobra.Command {
	// rootCmd represents the base command when called without any subcommands
	rootCmd := &cobra.Command{
		Use:   "lr",
		Short: "LR CLI",
		Long:  `LoginRadius CLI to support LoginRadius API's and workflow`,
	}

	// Authentication Commands
	loginCmd := login.NewLoginCmd()
	rootCmd.AddCommand((loginCmd))

	return rootCmd
}
