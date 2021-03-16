package verify

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/loginradius/lr-cli/cmd/verify/resend"
	"github.com/loginradius/lr-cli/cmdutil"
	"github.com/loginradius/lr-cli/config"
	"github.com/loginradius/lr-cli/request"
	"github.com/spf13/cobra"
)

type VerifyOpts struct {
	Email    string `json:"Email"`
	Username string `json:"Username"`
}

type Result struct {
	IsExist bool `json:IsExist`
}

var url string

func NewVerifyCmd() *cobra.Command {
	opts := &VerifyOpts{}

	cmd := &cobra.Command{
		Use:   "verify",
		Short: "Verify Email/Password",
		Long:  `This commmand verfies if email/username exists on your site or not`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.Email == "" && opts.Username == "" {
				return &cmdutil.FlagError{Err: errors.New("`--email` or `--username` is require argument")}
			}
			if opts.Email != "" && opts.Username != "" {
				return &cmdutil.FlagError{Err: errors.New("Enter only `--email` or `--username` is require argument")}
			}
			return verify(opts)

		},
	}
	resendCmd := resend.NewResendCmd()
	cmd.AddCommand(resendCmd)

	fl := cmd.Flags()
	fl.StringVarP(&opts.Email, "email", "e", "", "Email Value")
	fl.StringVarP(&opts.Username, "username", "u", "", "Username value")

	return cmd
}

func verify(opts *VerifyOpts) error {
	conf := config.GetInstance()
	if opts.Username != "" && opts.Email == "" {
		url = conf.LoginRadiusAPIDomain + "/identity/v2/auth/username?apikey=" + conf.LoginRadiusAPIKey + "&username=" + opts.Username
	} else if opts.Email != "" && opts.Username == "" {
		url = conf.LoginRadiusAPIDomain + "/identity/v2/auth/email?apikey=" + conf.LoginRadiusAPIKey + "&email=" + opts.Email
	} else {
		fmt.Println("Use paramters correctly")
	}
	var resultResp Result
	resp, err := request.Rest(http.MethodGet, url, nil, "")
	err = json.Unmarshal(resp, &resultResp)
	if err != nil {
		return err
	}
	fmt.Println(resultResp.IsExist)
	return nil
}
