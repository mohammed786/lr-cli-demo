package logout

import (
	"errors"
	"io/ioutil"
	"os"
	"os/user"
	"path"
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
	dirName := filepath.Join(user.HomeDir, ".lrcli")
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return &cmdutil.FlagError{Err: errors.New(" You have already been logged Out")}
	} else {
		dir, err := ioutil.ReadDir(dirName)
		for _, d := range dir {
			os.RemoveAll(path.Join([]string{dirName, d.Name()}...))
		}
		if err != nil {
			return err
		}

	}
	return nil
}
