package domain

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

type domainManagement struct {
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

	cmd := &cobra.Command{
		Use:   "domain",
		Short: "get domains added",
		Long:  `This commmand lists domains added`,
		RunE: func(cmd *cobra.Command, args []string) error {

			return get()

		},
	}

	return cmd
}

func get() error {
	conf := config.GetInstance()

	url = conf.LoginRadiusAPIDomain + "/deployment/sites?"

	v2, err := getCreds()
	headersV := map[string]string{
		"Origin":                  "https://dev-dashboard.lrinternal.com",
		"x-is-loginradius--sign":  v2.XSign,
		"x-is-loginradius--token": v2.XToken,
		"x-is-loginradius-ajax":   "true",
	}

	var resultResp domainManagement
	resp, err := request.Rest(http.MethodGet, url, headersV, "")
	err = json.Unmarshal(resp, &resultResp)
	if err != nil {
		return err
	}
	fmt.Println(resultResp)
	return nil
}
