package social

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"lr-cli/cmd/add/social"
	"lr-cli/cmdutil"
	"lr-cli/config"
	"lr-cli/request"
	"net/http"
	"os"
	"os/user"
	"path/filepath"

	"github.com/spf13/cobra"
)

type LoginResponse struct {
	XSign  string `json:"xsign"`
	XToken string `json:"xtoken"`
}

var fileName string

type provider struct {
	Provider string `json:"provider"`
}

type Result struct {
	Isdeleted bool `json:"isdeleted"`
}

var url string

func getCreds() (*LoginResponse, error) {
	var v2 LoginResponse
	user, _ := user.Current()
	fileName = filepath.Join(user.HomeDir, ".lrcli", "token.json")
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return (&v2), (err)
	}

	file, _ := ioutil.ReadFile(fileName)
	json.Unmarshal(file, &v2)
	return (&v2), (err)
}

func NewsocialCmd() *cobra.Command {
	opts := &provider{}

	cmd := &cobra.Command{
		Use:   "social",
		Short: "delete social provider",
		Long:  `This commmand deletes social provider`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.Provider == "" {
				return &cmdutil.FlagError{Err: errors.New("`provider` is require argument")}
			}
			return delete(opts)

		},
	}
	socialCmd := social.NewsocialCmd()
	cmd.AddCommand(socialCmd)

	fl := cmd.Flags()
	fl.StringVarP(&opts.Provider, "provider", "p", "", "provider name")

	return cmd
}

func delete(opts *provider) error {
	conf := config.GetInstance()
	if opts.Provider != "" {
		url = conf.LoginRadiusAPIDomain + "/platform-configuration/social-provider-config-remove?"
	} else {
		fmt.Println("Use paramters correctly")
	}
	v2, err := getCreds()
	headersV := map[string]string{
		"Origin":                  "https://dev-dashboard.lrinternal.com",
		"x-is-loginradius--sign":  v2.XSign,
		"x-is-loginradius--token": v2.XToken,
		"x-is-loginradius-ajax":   "true",
	}

	var resultResp Result
	resp, err := request.Rest(http.MethodDelete, url, headersV, opts.Provider)
	err = json.Unmarshal(resp, &resultResp)
	if err != nil {
		return err
	}
	fmt.Println(resultResp.Isdeleted)
	return nil
}
