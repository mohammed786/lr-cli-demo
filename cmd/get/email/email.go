package email

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/loginradius/lr-cli/config"
	"github.com/loginradius/lr-cli/request"

	"github.com/spf13/cobra"
)

var fileName string

type email struct {
	Description string `json:"description"`
	ErrorCode   string `json:"errorCode"`
	Message     string `json:"message"`
}

var url string

func NewemailCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "email",
		Short:   "get email config",
		Long:    `This commmand lists email config`,
		Example: `$ lr get email`,
		RunE: func(cmd *cobra.Command, args []string) error {

			return get()

		},
	}

	return cmd
}

func get() error {
	conf := config.GetInstance()

	url = conf.LoginRadiusAPIDomain + "/platform-configuration/global-email-configuration?"

	var resultResp email
	resp, err := request.Rest(http.MethodGet, url, nil, "")
	err = json.Unmarshal(resp, &resultResp)
	if err != nil {
		return err
	}
	fmt.Println(resultResp)
	return nil
}
