package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var JID string

// exactCmd represents the exact command
var exactCmd = &cobra.Command{
	Use:   "exact",
	Short: "Get Specific Joke",
	Long:  `Command that outputs specific joke. Requires a Joke ID as flag `,
	Run: func(cmd *cobra.Command, args []string) {
		getRandomJoker()
	},
}

func init() {
	rootCmd.AddCommand(exactCmd)
	exactCmd.PersistentFlags().StringVarP(&JID, "jokeID", "j", "", "Passes in Joke ID") //added a flag here

}

type Joker struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func getRandomJoker() {
	url := "https://icanhazdadjoke.com/j/" + JID
	responseBytes := getJokerData(url)
	joke := Joker{}
	if err := json.Unmarshal(responseBytes, &joke); err != nil {
		fmt.Printf("Could not unmarshal response % -v", err)
	}
	fmt.Println(string(joke.Joke))
}

func getJokerData(baseAPI string) []byte {
	request, err := http.NewRequest(
		http.MethodGet,
		baseAPI,
		nil,
	)
	if err != nil {
		log.Printf("Could not make request for a dad joke % -v", err)
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "Dadjoke CLI (https://github.com/spf13/cli_lr)")

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
