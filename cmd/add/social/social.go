package social

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/loginradius/lr-cli/cmd/get/social"
	"github.com/loginradius/lr-cli/cmdutil"
	"github.com/loginradius/lr-cli/config"
	"github.com/loginradius/lr-cli/request"

	"github.com/spf13/cobra"
)

var fileName string

type Provider struct {
	Provider string `json:"provider"`
}

type Result struct {
	Status bool `json:"status"`
}

var url1 string
var url2 string

func NewsocialCmd() *cobra.Command {
	opts := &Provider{}

	cmd := &cobra.Command{
		Use:     "social",
		Short:   "add social provider",
		Long:    `This commmand adds social provider`,
		Example: `$ lr add social --provider <provider>`,
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

	url1 = conf.LoginRadiusAPIDomain + "/platform-configuration/social-providers/status?"

	url2 = conf.LoginRadiusAPIDomain + "/platform-configuration/social-provider/options?"

	var resultResp Result
	resp1, err := request.Rest(http.MethodPost, url1, nil, opts.Provider)
	err = json.Unmarshal(resp1, &resultResp)
	if err != nil {
		return err
	}
	fmt.Println(resultResp.Status)

	var resultResp2 Result
	resp, err1 := request.Rest(http.MethodPost, url2, nil, opts.Provider)
	err1 = json.Unmarshal(resp, &resultResp2)
	if err1 != nil {
		return err
	}
	fmt.Println(resultResp2.Status)
	return nil
}
