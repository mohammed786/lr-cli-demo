package request

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/loginradius/lr-cli/cmdutil"
)

func Rest(method string, url string, headers map[string]string, payload string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, strings.NewReader(payload))

	if err != nil {
		log.Printf("error while Performing the Request: %s", err.Error())
		return nil, err
	}

	// LoginRadius Default Headers
	v2, err := cmdutil.GetCreds()
	if err == nil && v2.AppName != "" {
		req.Header.Set("x-is-loginradius--sign", v2.XSign)
		req.Header.Set("x-is-loginradius--token", v2.XToken)
	}
	req.Header.Set("Origin", "https://dev-dashboard.lrinternal.com")
	req.Header.Set("x-is-loginradius-ajax", "true")

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	if headers != nil {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}

	resp, err := client.Do(req)

	if err != nil {
		log.Printf("%s", err.Error())
	}

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
