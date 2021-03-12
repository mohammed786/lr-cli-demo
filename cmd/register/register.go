package register

import (
	"fmt"

	"github.com/pkg/browser"
	"github.com/spf13/cobra"
)

var appName, url string
var err error

type app struct {
	IsSiteRegistered bool `json: IsSiteRegistered `
}

func NewRegisterCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "register",
		Short: "Register a LR account",
		Long:  `This commmand registers a user to a LR account`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return register()
		},
	}
	return cmd
}

func register() error {
	fmt.Println("Registration tab will open in browser, Press enter to continue")
	fmt.Scanln()
	browser.OpenURL("https://accounts.loginradius.com/auth.aspx?action=register&return_url=https://dashboard.loginradius.com/login")
	return nil
}
