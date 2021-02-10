package cmd

import (
	"github.com/spf13/cobra"
)

var username string //global vars
var email string
var url string

func init() {
	rootCmd.AddCommand(authCmd)
	authCmd.PersistentFlags().StringVarP(&username, "username", "u", "", "Username value") //flags
	authCmd.PersistentFlags().StringVarP(&email, "email", "e", "", "Email value")
}

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "For Auth API's",
	Long:  `Supports Loginradius Auth API's`,
}
