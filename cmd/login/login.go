package login

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"lr-cli/cmdutil"
	"lr-cli/config"
	"lr-cli/request"
	"net/http"
	"os"
	"os/user"
	"path/filepath"

	"github.com/spf13/cobra"
)

type LoginOpts struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

type Token struct {
	AccessToken string `json:"access_token"`
}

type LoginResponse struct {
	XSign   string `json:"xsign"`
	XToken  string `json:"xtoken"`
	AppName string `json:"app_name"`
}

var fileName string

func NewLoginCmd() *cobra.Command {

	opts := &LoginOpts{}
	cmd := &cobra.Command{
		Use:   "login",
		Short: "Login to LR account",
		Long:  `This commmand logs user into the LR account`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.Email == "" {
				return &cmdutil.FlagError{Err: errors.New("`--email` is require argument")}
			}
			if opts.Password == "" {
				return &cmdutil.FlagError{Err: errors.New("`--password` is require argument")}
			}
			return validateLogin(opts)

		},
	}
	fl := cmd.Flags()
	fl.StringVarP(&opts.Email, "email", "e", "", "Email Value")
	fl.StringVarP(&opts.Password, "password", "p", "", "Password value")

	return cmd
}

func validateLogin(opts *LoginOpts) error {
	var v1 Token     //access
	v2 := getCreds() //x
	if v2.XSign != "" && v2.XToken != "" {
		conf := config.GetInstance()
		validateURL := conf.AdminConsoleAPIDomain + "/auth/validatetoken"
		headersV := map[string]string{
			"Origin":                  "https://dev-dashboard.lrinternal.com",
			"x-is-loginradius--sign":  v2.XSign,
			"x-is-loginradius--token": v2.XToken,
			"x-is-loginradius-ajax":   "true",
		}

		resp, err := request.Rest(http.MethodGet, validateURL, headersV, "")
		if err != nil {
			return err
		}
		err = json.Unmarshal(resp, &v1)
		if v1.AccessToken != "" {
			return &cmdutil.FlagError{Err: errors.New("Already logged in ")}
		} else {
			return doLogin(opts)
		}
	}
	return doLogin(opts)
}

func doLogin(opts *LoginOpts) error {
	log.Printf("%s", opts.Email)
	conf := config.GetInstance()
	hubPageURL := conf.LoginRadiusAPIDomain + "/identity/v2/auth/login/2FA?apiKey=" + conf.LoginRadiusAPIKey
	body, _ := json.Marshal(opts)
	resp, err := request.Rest(http.MethodPost, hubPageURL, nil, string(body))
	if err != nil {
		return err
	}

	// Identity API
	var identityResp Token
	err = json.Unmarshal(resp, &identityResp)
	if err != nil {
		return err
	}

	if identityResp.AccessToken == "" {
		return errors.New("Unable to get the Access token")
	}

	// Admin Console Backend API
	var resObj LoginResponse

	backendURL := conf.AdminConsoleAPIDomain + "/auth/login"
	body, _ = json.Marshal(map[string]string{
		"accesstoken": identityResp.AccessToken,
	})
	headers := map[string]string{
		"Origin":                "https://dev-dashboard.lrinternal.com",
		"x-is-loginradius-ajax": "true",
	}
	resp, err = request.Rest(http.MethodPost, backendURL, headers, string(body))

	err = json.Unmarshal(resp, &resObj)
	if err != nil {
		return err
	}
	return storeCreds(&resObj)
}

func storeCreds(cred *LoginResponse) error {
	user, _ := user.Current()

	os.Mkdir(filepath.Join(user.HomeDir, ".lrcli"), 0755)
	fileName = filepath.Join(user.HomeDir, ".lrcli", "token.json")

	dataBytes, _ := json.Marshal(cred)

	return ioutil.WriteFile(fileName, dataBytes, 0644)

}

func getCreds() *LoginResponse {
	var v2 LoginResponse
	user, _ := user.Current()
	fileName = filepath.Join(user.HomeDir, ".lrcli", "token.json")
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		log.Println("Creating token.json")
	}

	file, _ := ioutil.ReadFile(fileName)
	json.Unmarshal(file, &v2)
	return &v2
}
