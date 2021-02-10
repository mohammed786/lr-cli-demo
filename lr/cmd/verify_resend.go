package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

func init() {
	authCmd.AddCommand(verify_resend)
}

//var url string
var verify_resend = &cobra.Command{
	Use:     "verify_resend",
	Aliases: []string{"verifyResend", "vr"},
	Short:   "Resends verification mail to email ID",
	Long:    `This command resends verification email to the entered email ID`,
	Run: func(cmd *cobra.Command, args []string) {
		getData1()
	},
}

type Verify struct { //for response
	IsPosted bool `json:IsPosted`
}

type Body struct { //for request
	Email string `json:Email`
}

func getData1() {
	if email != "" {
		url = "https://api.loginradius.com/identity/v2/auth/register?apikey=b9d66d19-e103-4de7-911f-f7d654fe0e3d&verificationurl=&emailtemplate="
	}
	responseBytes := gethttp1(url)
	verify := Verify{}
	err := json.Unmarshal(responseBytes, &verify)
	if err != nil {
		fmt.Printf("Could not unmarshal response % -v", err)
	}
	fmt.Println(verify.IsPosted)
}

func gethttp1(baseAPI string) []byte {
	body := Body{
		Email: email, //Post Body
	}
	json, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	request, err := http.NewRequest(
		http.MethodPut,
		baseAPI,
		bytes.NewBuffer(json),
	)
	if err != nil {
		log.Printf("Could not make request % -v", err)
	}

	request.Header.Add("content-Type", "application/json")

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
