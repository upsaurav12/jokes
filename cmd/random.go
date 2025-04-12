/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		getRandomJoke()
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)
}

type Joke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func getRandomJoke() string {
	url := "https://icanhazdadjoke.com"
	responseByte := getJokeData(url)

	joke := &Joke{}
	if err := json.Unmarshal(responseByte, joke); err != nil {
		log.Printf("Could not unmarshal respnse %v", err)
	}

	// fmt.Println(string(joke.Joke))

	return string(joke.Joke)
}

func getJokeData(baseurl string) []byte {
	request, err := http.NewRequest(
		http.MethodGet,
		baseurl,
		nil,
	)

	if err != nil {
		log.Printf("Could not request a dadjoke %v", err)
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "Dadjoke CLI (github.com/example)")

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		log.Printf("Could not get response for dadjoke %v", err)
	}

	responseByte, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Printf("Could not get response in Bytes %v", err)
	}

	return responseByte
}
