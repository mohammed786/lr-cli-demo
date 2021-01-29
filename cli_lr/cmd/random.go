package cmd

import (
	"encoding/json"
	"fmt"

	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Get random dad joke",
	Long:  `Gets a random dad joke from the icanhazdadjoke api`,
	Run: func(cmd *cobra.Command, args []string) {
		getRandomJoke()
	},
}

func init() {
	rootCmd.AddCommand(randomCmd) //here is where we can declare config and flags

}

type Joke struct {
	ID     string `json:"id"` //models the response pattern
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func getRandomJoke() {
	url := "https://icanhazdadjoke.com/"
	responseBytes := getJokeData(url)
	joke := Joke{}
	if err := json.Unmarshal(responseBytes, &joke); err != nil {
		fmt.Printf("Could not unmarshal response % -v", err)
	} //stores the result in interface
	fmt.Println(string(joke.Joke))
}

func getJokeData(baseAPI string) []byte {
	request, err := http.NewRequest( //for the request
		http.MethodGet,
		baseAPI,
		nil,
	)
	if err != nil {
		log.Printf("Could not make request for a dad joke % -v", err)
	}

	request.Header.Add("Accept", "application/json") //custom headers
	request.Header.Add("User-Agent", "Dadjoke CLI (https://github.com/spf13/cli_lr)")

	response, err := http.DefaultClient.Do(request) //creates reponse

	if err != nil {
		log.Printf("Could not get the response % -v", err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body) //reads the body of the response
	if err != nil {
		log.Printf("Could not read the response% -v", err)
	}

	return responseBytes
}
