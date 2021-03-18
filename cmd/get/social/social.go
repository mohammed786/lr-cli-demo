package social

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/loginradius/lr-cli/config"
	"github.com/loginradius/lr-cli/request"
	"github.com/spf13/cobra"
)

var fileName string

type socialProvider struct {
	Name    string `json:"name"`
	Display string `json:"display"`
	Mdfile  string `json:"mdfile"`
}

type socialProviderList struct {
	Data []socialProvider `json:"data"`
}

var url string

func NewsocialCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "social",
		Short:   "get social providers",
		Long:    `This commmand lists social providers`,
		Example: `$ lr get social`,
		RunE: func(cmd *cobra.Command, args []string) error {

			return get()

		},
	}

	return cmd
}

func get() error {
	conf := config.GetInstance()

	url = conf.LoginRadiusAPIDomain + "/platform-configuration/social-provider/list?"

	var resultResp socialProviderList
	resp, err := request.Rest(http.MethodGet, url, nil, "")
	err = json.Unmarshal(resp, &resultResp)
	if err != nil {
		return err
	}
	fmt.Println(resultResp)
	return nil
}
