package email

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

type email struct {
	EmailLinkExpire            int `json:"EmailLinkExpire"`
	EmailNotificationCount     int `json:"EmailNotificationCount"`
	EmailNotificationFrequency int `json:"EmailNotificationFrequency"`
}

var url string

func NewemailCmd() *cobra.Command {
	opts := &email{}
	cmd := &cobra.Command{
		Use:     "email",
		Short:   "set email config",
		Long:    `This commmand sets email config`,
		Example: `$ lr set email`,
		RunE: func(cmd *cobra.Command, args []string) error {
			var intpointer1 *int
			intpointer1 = &opts.EmailLinkExpire
			var intpointer2 *int
			intpointer2 = &opts.EmailNotificationCount
			var intpointer3 *int
			intpointer3 = &opts.EmailNotificationFrequency

			if intpointer1 == nil {
				return &cmdutil.FlagError{Err: errors.New("`email_link_expire` is require argument")}
			}

			if intpointer2 == nil {
				return &cmdutil.FlagError{Err: errors.New("`email_notif_count` is require argument")}
			}

			if intpointer3 == nil {
				return &cmdutil.FlagError{Err: errors.New("`email_notif_frequency` is require argument")}
			}
			//fmt.Printf(string(rune(opts.EmailLinkExpire)))
			return set(opts.EmailLinkExpire, opts.EmailNotificationCount, opts.EmailNotificationFrequency)

		},
	}
	fl := cmd.Flags()

	fl.IntVarP(&opts.EmailLinkExpire, "email_link_expire", "a", 10080, "usage")
	fl.IntVarP(&opts.EmailNotificationCount, "email_notif_count", "b", 3, "usage")
	fl.IntVarP(&opts.EmailNotificationFrequency, "email_notif_frequency", "c", 1440, "domain name")

	return cmd
}

func set(a int, b int, c int) error {
	conf := config.GetInstance()

	url = conf.AdminConsoleAPIDomain + "/platform-configuration/global-email-configuration?"

	body, _ := json.Marshal(map[string]int{
		"EmailLinkExpire":            a,
		"EmailNotificationCount":     b,
		"EmailNotificationFrequency": c,
	})

	var resultResp email
	resp, err := request.Rest(http.MethodPut, url, nil, string(body))
	err = json.Unmarshal(resp, &resultResp)
	if err != nil {
		return err
	}
	//fmt.Printf("%s", string(resp))
	fmt.Printf("{EmailLinkExpire, EmailNotificationCount, EmailNotificationFrequency}\n")
	//fmt.Printf(string(resp))
	fmt.Println(resultResp)
	return nil
}
