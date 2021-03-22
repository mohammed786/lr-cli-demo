package add

import (
	"github.com/loginradius/lr-cli/cmd/add/domain"
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

	return cmd
}
