package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/pkg/browser"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var emailLogin string
var password string
var url1 string
var accessToken string

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to LR account",
	Long:  `This commmand logs user into the LR account`,
	Run: func(cmd *cobra.Command, args []string) {
		getData2()
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.PersistentFlags().StringVarP(&emailLogin, "email", "e", "", "Email value")
	loginCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "Password value")
}

func prompts() {
	fmt.Println("To authenticate, Press Enter to open loginradius.com in Browser") //alt method
	fmt.Scanln()
	browser.OpenURL("https://accounts.loginradius.com/auth.aspx?return_url=https://dashboard.loginradius.com/login")
}

type Body1 struct { //request
	Email    string `json:Email`
	Password string `json:Password`
}
type Token struct {
	Accesstoken string `json:accesstoken`
}

type LoginResponse struct { //response
	XSign  string `json:"xsign"`
	XToken string `json:"xtoken"`
}

func getData2() {
	if emailLogin != "" && password != "" {
		at := make(map[string]interface{})
		url1 = "https://devapi.lrinternal.com/identity/v2/auth/login/2FA?apiKey=ddff8a63-cbc3-4723-8415-b910c4d8770d&loginUrl=&emailTemplate=&verificationUrl=https://demotesting.devhub.lrinternal.com/auth.aspx?return_url=https://dev-dashboard.lrinternal.com/login&smsTemplate=&smsTemplate2FA="
		responseBytes := gethttp2(url1)
		err := json.Unmarshal(responseBytes, &at)
		if err != nil {
			fmt.Printf("Could not unmarshal response % -v", err)
		}
		//filename := filepath.Join("/Users", "akashpatil", ".lrcli", "accessTokens.json")
		filename := "myFile.json"
		errFile := checkFile(filename)
		if errFile != nil {
			logrus.Error(err)
		}

		file, err := ioutil.ReadFile(filename)
		if err != nil {
			logrus.Error(err)
		}

		accessToken = at["access_token"].(string)
		if accessToken != "" {
			url2 := "https://devadmin-console-api.lrinternal.com/auth/login?appName=&"

			responseBytes1 := gethttp3(url2)

			if err != nil {
				fmt.Printf("Could not unmarshal response % -v", err)
			}
			var resObj LoginResponse
			json.Unmarshal(responseBytes1, &resObj)
			x_token := resObj.XToken
			x_sign := resObj.XSign
			data := []LoginResponse{}

			json.Unmarshal(file, &data)

			newStruct := &LoginResponse{
				XToken: x_token,
				XSign:  x_sign,
			}
			data = append(data, *newStruct)

			dataBytes, err := json.Marshal(data[len(data)-1])

			err = ioutil.WriteFile(filename, dataBytes, 0644)
			if err != nil {
				logrus.Error(err)
			}
		} else {
			fmt.Println("Need access token")
		}

	} else {
		prompts()
	}
}

//Used for getting access token
func gethttp2(baseAPI string) []byte {
	body1 := Body1{
		Email:    emailLogin,
		Password: password,
	}
	json, err := json.Marshal(body1) //struct to byte
	if err != nil {
		panic(err)
	}
	request, err := http.NewRequest(
		http.MethodPost,
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

//Used for getting XToken and XSign
func gethttp3(API string) []byte {
	client := &http.Client{}
	postBody, _ := json.Marshal(map[string]string{
		"accesstoken": accessToken,
	})
	responseBody := bytes.NewBuffer(postBody)
	req, err := http.NewRequest(http.MethodPost, API, responseBody)
	if err != nil {
		log.Printf("%s", err.Error())
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Origin", "https://dev-dashboard.lrinternal.com")
	req.Header.Add("x-is-loginradius-ajax", "true")

	resp, err := client.Do(req)

	if err != nil {
		log.Printf("%s", err.Error())
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	return bodyBytes
}

func checkFile(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			return err
		}
	}
	return nil
}
