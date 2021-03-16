package config

import "github.com/spf13/cobra"

func NewConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Shows/Stores App's API Key/Secret",
		Long:  `This command displays and stores the User App's API Key/Secret`,
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	return cmd
}
