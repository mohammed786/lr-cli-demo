package domain

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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

type domain struct {
	// CallbackUrl string `json:"CallbackUrl"`
	Domain     string `json:"http://localhost"`
	Production string `json:"production"`
	staging    string `json:""`
}

type Result struct {
	CallbackUrl string `json:"CallbackUrl"`
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

func NewdomainCmd() *cobra.Command {
	opts := &domain{}

	cmd := &cobra.Command{
		Use:   "domain",
		Short: "add doamin",
		Long:  `This commmand adds domain`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.Production == "" {
				return &cmdutil.FlagError{Err: errors.New("`production` is require argument")}
			}
			return add(opts)

		},
	}
	fl := cmd.Flags()
	fl.StringVarP(&opts.Production, "production", "p", "", "production name")

	return cmd
}

func add(opts *domain) error {
	body, _ := json.Marshal(opts)
	conf := config.GetInstance()
	if opts.Production != "" {
		url = conf.LoginRadiusAPIDomain + "deployment/sites?"
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
	resp, err := request.Rest(http.MethodPost, url, headersV, string(body))
	err = json.Unmarshal(resp, &resultResp)
	if err != nil {
		return err
	}
	fmt.Println(resultResp.CallbackUrl)
	return nil
}
