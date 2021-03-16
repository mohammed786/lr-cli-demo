package logout

import (
	"errors"
	"os"
	"os/user"
	"path/filepath"

	"github.com/loginradius/lr-cli/cmdutil"
	"github.com/spf13/cobra"
)

func NewLogoutCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "logout",
		Short: "Logout of LR account",
		Long:  `This commmand logs user out of the LR account`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return logout()
		},
	}
	return cmd
}

func logout() error {
	user, _ := user.Current()
	fileName := filepath.Join(user.HomeDir, ".lrcli", "token.json")
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return &cmdutil.FlagError{Err: errors.New(" You have already been logged Out")}
	} else {
		err = os.Remove(fileName)
		if err != nil {
			return err
		}
	}
	return nil
}
