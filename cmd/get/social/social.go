package social

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

type socialProvider struct {
	Name    string `json:"name"`
	Display string `json:"display"`
	Mdfile  string `json:"mdfile"`
}

type socialProviderList struct {
	Data []socialProvider `json:"data"`
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

	cmd := &cobra.Command{
		Use:   "social",
		Short: "get social providers",
		Long:  `This commmand lists social providers`,
		RunE: func(cmd *cobra.Command, args []string) error {

			return get()

		},
	}

	return cmd
}

func get() error {
	conf := config.GetInstance()

	url = conf.LoginRadiusAPIDomain + "/platform-configuration/social-provider/list?"

	v2, err := getCreds()
	headersV := map[string]string{
		"Origin":                  "https://dev-dashboard.lrinternal.com",
		"x-is-loginradius--sign":  v2.XSign,
		"x-is-loginradius--token": v2.XToken,
		"x-is-loginradius-ajax":   "true",
	}

	var resultResp socialProviderList
	resp, err := request.Rest(http.MethodGet, url, headersV, "")
	err = json.Unmarshal(resp, &resultResp)
	if err != nil {
		return err
	}
	fmt.Println(resultResp)
	return nil
}
