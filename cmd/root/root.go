package root

import (
	"github.com/loginradius/lr-cli/cmd/get"
	"github.com/loginradius/lr-cli/cmd/login"
	"github.com/loginradius/lr-cli/cmd/logout"
	"github.com/loginradius/lr-cli/cmd/register"
	"github.com/loginradius/lr-cli/cmd/resetSecret"
	"github.com/loginradius/lr-cli/cmd/verify"
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

	logoutCmd := logout.NewLogoutCmd()
	rootCmd.AddCommand((logoutCmd))

	registerCmd := register.NewRegisterCmd()
	rootCmd.AddCommand((registerCmd))

	verifyCmd := verify.NewVerifyCmd()
	rootCmd.AddCommand((verifyCmd))

	getCmd := get.NewGetCmd()
	rootCmd.AddCommand((getCmd))

	resetCmd := resetSecret.NewResetCmd()
	rootCmd.AddCommand((resetCmd))

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.lr-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	return rootCmd
}
