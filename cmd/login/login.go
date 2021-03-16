package login

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/MakeNowJust/heredoc"

	"github.com/loginradius/lr-cli/cmdutil"
	"github.com/loginradius/lr-cli/config"
	"github.com/loginradius/lr-cli/request"
	"github.com/spf13/cobra"
)

type LoginOpts struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

func NewLoginCmd() *cobra.Command {

	opts := &LoginOpts{}
	cmd := &cobra.Command{
		Use:   "login",
		Short: "Login to LR account",
		Long: heredoc.Doc(`
			This commmand logs user into the LR account.
		`),
		Example: heredoc.Doc(`
			$ lr login -e <email> -p <password>
		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.Email == "" {
				return &cmdutil.FlagError{Err: errors.New("`--email` is require argument")}
			}
			if opts.Password == "" {
				return &cmdutil.FlagError{Err: errors.New("`--password` is require argument")}
			}
			isValid, err := validateLogin()

			if err != nil {
				return err
			} else if isValid {
				log.Printf("%s", "You are already been logged in")
				return nil
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
	var identityResp cmdutil.Token
	err = json.Unmarshal(resp, &identityResp)
	if err != nil {
		return err
	}

	if identityResp.AccessToken == "" {
		return errors.New("Unable to get the Access token")
	}

	// Admin Console Backend API
	var resObj cmdutil.LoginResponse

	backendURL := conf.AdminConsoleAPIDomain + "/auth/login"
	body, _ = json.Marshal(map[string]string{
		"accesstoken": identityResp.AccessToken,
	})
	resp, err = request.Rest(http.MethodPost, backendURL, nil, string(body))

	err = json.Unmarshal(resp, &resObj)
	if err != nil {
		return err
	}
	log.Println("Successfully Logged In")
	return cmdutil.StoreCreds(&resObj)
}

func validateLogin() (bool, error) {
	conf := config.GetInstance()
	validateURL := conf.AdminConsoleAPIDomain + "/auth/validatetoken"
	resp, err := request.Rest(http.MethodGet, validateURL, nil, "")
	if err != nil {
		return false, err
	}
	var v1 cmdutil.Token
	err = json.Unmarshal(resp, &v1)
	if v1.AccessToken != "" {
		return true, nil
	}
	return false, nil
}
