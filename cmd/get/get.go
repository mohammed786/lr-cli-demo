package get

import (
	"github.com/loginradius/lr-cli/cmd/get/config"
	"github.com/loginradius/lr-cli/cmd/get/social"
	"github.com/loginradius/lr-cli/cmd/get/theme"
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

	socialCmd := social.NewsocialCmd()
	cmd.AddCommand(socialCmd)

	configCmd := config.NewConfigCmd()
	cmd.AddCommand(configCmd)
	return cmd
}
