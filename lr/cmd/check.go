package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

func init() {
	authCmd.AddCommand(check)
}

var check = &cobra.Command{
	Use:   "check",
	Short: "Checks is Email/UserName exists on your site",
	Long:  `This command checks if Email/UserName exists on your site or not`,
	Run: func(cmd *cobra.Command, args []string) {
		getData()
	},
}

type Check struct {
	IsExist bool `json:IsExist`
}

func getData() {

	if username != "" && email == "" {
		url = "https://api.loginradius.com/identity/v2/auth/username?apikey=b9d66d19-e103-4de7-911f-f7d654fe0e3d&username=" + username
	} else if email != "" && username == "" {
		url = "https://api.loginradius.com/identity/v2/auth/email?apikey=b9d66d19-e103-4de7-911f-f7d654fe0e3d&email=" + email
	} else {
		fmt.Println("Use paramters correctly")
	}

	responseBytes := gethttp(url)
	check := Check{}
	err := json.Unmarshal(responseBytes, &check)
	if err != nil {
		fmt.Printf("Could not unmarshal response % -v", err)
	}
	fmt.Println(check.IsExist)

}

func gethttp(baseAPI string) []byte {
	request, err := http.NewRequest(
		http.MethodGet,
		baseAPI,
		nil,
	)
	if err != nil {
		log.Printf("Could not make request % -v", err)
	}

	request.Header.Add("Accept", "application/json")

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		log.Printf("Could not get the response % -v", err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Could not read the response% -v", err)
	}

	return responseBytes
}
