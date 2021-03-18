package domain

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/loginradius/lr-cli/config"
	"github.com/loginradius/lr-cli/request"
	"github.com/spf13/cobra"
)

type LoginResponse struct {
	XSign  string `json:"xsign"`
	XToken string `json:"xtoken"`
}

var fileName string

type domainManagement struct {
	CallbackUrl string `json:"CallbackUrl"`
}

var url string

func NewdomainCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "domain",
		Short:   "get domains added",
		Long:    `This commmand lists domains added`,
		Example: `$ lr get domain`,
		RunE: func(cmd *cobra.Command, args []string) error {

			return get()

		},
	}

	return cmd
}

func get() error {
	conf := config.GetInstance()

	url = conf.LoginRadiusAPIDomain + "/deployment/sites?"

	var resultResp domainManagement
	resp, err := request.Rest(http.MethodGet, url, nil, "")
	err = json.Unmarshal(resp, &resultResp)
	if err != nil {
		return err
	}
	fmt.Println(resultResp)
	return nil
}
