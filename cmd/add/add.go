package add

import (
	"lr-cli/cmd/add/domain"
	"lr-cli/cmd/add/social"

	"github.com/spf13/cobra"
)

func NewaddCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "add",
		Short: "add command",
		Long:  `This commmand acts as a base command for add subcommands`,
	}

	domainCmd := domain.NewdomainCmd()
	cmd.AddCommand(domainCmd)

	socialCmd := social.NewsocialCmd()
	cmd.AddCommand(socialCmd)
	return cmd
}
