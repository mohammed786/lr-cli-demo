package social

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"lr-cli/cmd/get/social"
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

type Provider struct {
	Provider string `json:"provider"`
}

type Result struct {
	Status bool `json:"status"`
}

var url1 string
var url2 string

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
	opts := &Provider{}

	cmd := &cobra.Command{
		Use:   "social",
		Short: "add social provider",
		Long:  `This commmand adds social provider`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.Provider == "" {
				return &cmdutil.FlagError{Err: errors.New("`provider` is require argument")}
			}
			return add(opts)

		},
	}
	socialCmd := social.NewsocialCmd()
	cmd.AddCommand(socialCmd)

	fl := cmd.Flags()
	fl.StringVarP(&opts.Provider, "provider", "p", "", "provider name")

	return cmd
}

func add(opts *Provider) error {
	conf := config.GetInstance()
	if opts.Provider != "" {
		url1 = conf.LoginRadiusAPIDomain + "/platform-configuration/social-providers/status?"
	} else {
		fmt.Println("Use paramters correctly")
	}

	if opts.Provider != "" {
		url2 = conf.LoginRadiusAPIDomain + "/platform-configuration/social-provider/options?"
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
	resp1, err := request.Rest(http.MethodPost, url1, headersV, opts.Provider)
	err = json.Unmarshal(resp1, &resultResp)
	if err != nil {
		return err
	}
	fmt.Println(resultResp.Status)

	var resultResp2 Result
	resp, err1 := request.Rest(http.MethodPost, url2, headersV, opts.Provider)
	err1 = json.Unmarshal(resp, &resultResp2)
	if err1 != nil {
		return err
	}
	fmt.Println(resultResp2.Status)
	return nil
}
