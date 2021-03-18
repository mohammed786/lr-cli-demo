package social

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/loginradius/lr-cli/cmdutil"
	"github.com/loginradius/lr-cli/config"
	"github.com/loginradius/lr-cli/request"

	"github.com/spf13/cobra"
)

var fileName string

type provider struct {
	providerName string `json:"providerName"`
}

type Result struct {
	Isdeleted bool `json:"isdeleted"`
}

var url string

func NewsocialCmd() *cobra.Command {
	opts := &provider{}

	cmd := &cobra.Command{
		Use:     "social",
		Short:   "delete social provider",
		Long:    `This commmand deletes social provider`,
		Example: `$ lr delete social --provider <provider>`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.providerName == "" {
				return &cmdutil.FlagError{Err: errors.New("`provider` is require argument")}
			}
			return delete(opts)

		},
	}

	fl := cmd.Flags()
	fl.StringVarP(&opts.providerName, "provider", "p", "", "provider name")

	return cmd
}

func delete(opts *provider) error {
	conf := config.GetInstance()

	url = conf.AdminConsoleAPIDomain + "/platform-configuration/social-provider-config-remove?"

	var resultResp Result
	resp, err := request.Rest(http.MethodDelete, url, nil, opts.providerName)
	err = json.Unmarshal(resp, &resultResp)
	if err != nil {
		return err
	}
	fmt.Println(resultResp.Isdeleted)
	return nil
}
