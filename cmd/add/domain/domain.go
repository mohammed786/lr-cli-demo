package domain

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/loginradius/lr-cli/request"

	"github.com/loginradius/lr-cli/cmdutil"
	"github.com/loginradius/lr-cli/config"

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

func NewdomainCmd() *cobra.Command {
	opts := &domain{}

	cmd := &cobra.Command{
		Use:     "domain",
		Short:   "add doamin",
		Long:    `This commmand adds domain`,
		Example: `$ lr add domain --production <production>`,
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
	var url string
	body, _ := json.Marshal(opts)
	conf := config.GetInstance()

	url = conf.LoginRadiusAPIDomain + "deployment/sites?"

	var resultResp Result
	resp, err := request.Rest(http.MethodPost, url, nil, string(body))
	err = json.Unmarshal(resp, &resultResp)
	if err != nil {
		return err
	}
	fmt.Println(resultResp.CallbackUrl)
	return nil
}
