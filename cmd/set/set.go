package set

import (
	"github.com/loginradius/lr-cli/cmd/set/domain"

	"github.com/spf13/cobra"
)

func NewsetCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "set",
		Short: "set command",
		Long:  `This commmand acts as a base command for set subcommands`,
	}

	domainCmd := domain.NewdomainCmd()
	cmd.AddCommand((domainCmd))

	return cmd
}
