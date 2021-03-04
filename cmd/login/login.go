package login

import (
	"encoding/json"
	"errors"
	"log"
	"lr-cli/cmdutil"
	"lr-cli/config"
	"lr-cli/request"
	"net/http"

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
			return doLogin(opts)
		},
	}
	fl := cmd.Flags()
	fl.StringVarP(&opts.Email, "email", "e", "", "Email Value")
	fl.StringVarP(&opts.Password, "password", "p", "", "Password value")

	return cmd
}

func doLogin(opts *LoginOpts) error {
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

	log.Printf("%s", resObj.AppName)

	return nil
}
