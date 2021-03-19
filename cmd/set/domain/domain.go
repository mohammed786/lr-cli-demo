package domain

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/loginradius/lr-cli/request"

	"github.com/loginradius/lr-cli/cmdutil"
	"github.com/loginradius/lr-cli/config"

	"github.com/spf13/cobra"
)

var fileName string

type domainManagement struct {
	CallbackUrl string `json:"CallbackUrl"`
}

type domain struct {
	// CallbackUrl string `json:"CallbackUrl"`
	Domain    string `json:"domain"`
	DomainMod string `json:"domainmod"`
}

type Result struct {
	CallbackUrl string `json:"CallbackUrl"`
}

func NewdomainCmd() *cobra.Command {
	opts := &domain{}

	cmd := &cobra.Command{
		Use:     "domain",
		Short:   "set domain",
		Long:    `This commmand sets domain`,
		Example: `$ lr set domain --domain <domain> --domainmod <domainmodified>`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.Domain == "" {
				return &cmdutil.FlagError{Err: errors.New("`domain` is require argument")}
			}

			if opts.DomainMod == "" {
				return &cmdutil.FlagError{Err: errors.New("`domainmod` is require argument")}
			}

			var p, _ = get()
			domain := strings.ReplaceAll(p.CallbackUrl, (";" + opts.Domain), (";" + opts.DomainMod))
			return delete(domain)

			// var p, _ = get()
			// fmt.Printf(p.CallbackUrl)
			// s := strings.Split(p.CallbackUrl, ";")
			// if len(s) < 3 {
			// 	domain := p.CallbackUrl + ";" + opts.Domain

			// 	return delete(domain)
			// } else {
			// 	return &cmdutil.FlagError{Err: errors.New("more than 3 domains cannot be added in free plan")}
			// }

		},
	}

	fl := cmd.Flags()
	fl.StringVarP(&opts.Domain, "domain", "d", "", "domain name")
	fl.StringVarP(&opts.DomainMod, "domainmod", "m", "", "domain modified name")

	return cmd
}

func get() (*domainManagement, error) {
	conf := config.GetInstance()
	var url string
	url = conf.AdminConsoleAPIDomain + "/deployment/sites?"

	var resultResp *domainManagement
	resp, err := request.Rest(http.MethodGet, url, nil, "")
	err = json.Unmarshal(resp, &resultResp)
	if err != nil {
		return nil, err
	}

	//append domain with resultResp
	//fmt.Print(resultResp.CallbackUrl)
	return resultResp, nil
}

func delete(domain string) error {
	var url string
	fmt.Printf("domain=%s", domain)
	body, _ := json.Marshal(map[string]string{
		"domain":     "http://localhost",
		"production": domain,
		"staging":    "",
	})
	conf := config.GetInstance()

	url = conf.AdminConsoleAPIDomain + "/deployment/sites?"

	// var resultResp2 domainManagement
	// resp, err2 := request.Rest(http.MethodGet, url, nil, "")
	// err2 = json.Unmarshal(resp, &resultResp2)
	// if err2 != nil {
	// 	return err2
	// }

	var resultResp Result
	resp, err := request.Rest(http.MethodPost, url, nil, string(body))
	err = json.Unmarshal(resp, &resultResp)
	if err != nil {
		return err
	}
	fmt.Println(resultResp.CallbackUrl)
	return nil
}
