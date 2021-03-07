package cmd

import (
	"github.com/spf13/cobra"
)

// siteCmd represents the site command
var siteCmd = &cobra.Command{
	Use:   "site",
	Short: "Handles the site of the user",
	Long:  `This command is used to make changes to the site of the user`,
}

func init() {
	rootCmd.AddCommand(siteCmd)

}
