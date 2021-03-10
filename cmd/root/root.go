package root

import (
	"lr-cli/cmd/get"
	"lr-cli/cmd/login"
	"lr-cli/cmd/verify"

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

	verifyCmd := verify.NewVerifyCmd()
	rootCmd.AddCommand((verifyCmd))

	getCmd := get.NewGetCmd()
	rootCmd.AddCommand((getCmd))

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.lr-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	return rootCmd
}
