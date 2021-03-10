package get

import (
	"lr-cli/cmd/get/social"
	"lr-cli/cmd/get/theme"

	"github.com/spf13/cobra"
)

func NewGetCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "get",
		Short: "get command",
		Long:  `This commmand acts as a base command for get subcommands`,
	}

	themeCmd := theme.NewThemeCmd()
	cmd.AddCommand(themeCmd)
	return cmd

	socialCmd := social.NewsocialCmd()
	cmd.AddCommand(socialCmd)
	return cmd
}
