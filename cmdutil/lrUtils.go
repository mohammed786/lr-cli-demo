package cmdutil

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
)

type LoginResponse struct {
	XSign   string `json:"xsign"`
	XToken  string `json:"xtoken"`
	AppName string `json:"app_name"`
}

type Token struct {
	AccessToken string `json:"access_token"`
}

func StoreCreds(cred *LoginResponse) error {
	user, _ := user.Current()

	os.Mkdir(filepath.Join(user.HomeDir, ".lrcli"), 0755)
	fileName := filepath.Join(user.HomeDir, ".lrcli", "token.json")

	dataBytes, _ := json.Marshal(cred)

	return ioutil.WriteFile(fileName, dataBytes, 0644)

}
func GetCreds() (*LoginResponse, error) {
	var v2 LoginResponse
	user, _ := user.Current()
	fileName := filepath.Join(user.HomeDir, ".lrcli", "token.json")
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return nil, err
	}

	file, _ := ioutil.ReadFile(fileName)
	json.Unmarshal(file, &v2)
	return &v2, nil
}
