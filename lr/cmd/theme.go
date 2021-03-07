package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// themeCmd represents the theme command
var theme = &cobra.Command{
	Use:   "theme",
	Short: "Handles theme of the site",
	Long:  `This command handles the theme of the site`,
	Run: func(cmd *cobra.Command, args []string) {
		list()
	},
}

func init() {
	siteCmd.AddCommand(theme)
}

//Displays the themes
func list() {
	fmt.Println("Available Themes:")
	fmt.Println("1. Tokyo")
	fmt.Println("2. London")
	fmt.Println("3. Helsinki")
}
