package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type Variable struct {
	VariableType     string `json:"variable_type"`
	Key              string `json:"key"`
	Value            string `json:"value"`
	Protected        bool   `json:"protected"`
	Masked           bool   `json:"masked"`
	EnvironmentScope string `json:"environment_scope"`
}

func main() {
	host := os.Args[2]
	projectIdInput := os.Args[3]
	privateToken := os.Args[4]

	client := http.Client{}
	url := fmt.Sprintf("%s/api/v4/projects/%s/variables", host, projectIdInput)
	var variables []Variable = []Variable{}

	for {
		var body []byte
		var tempVariables []Variable
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatalln(err)
		}

		req.Header.Add("PRIVATE-TOKEN", privateToken)

		resp, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}

		links := strings.Split(resp.Header.Get("Link"), ",")

		containsNext := false
		var nextUrlRaw string

		for _, link := range links {
			nextLink := strings.Split(link, ";")

			if strings.Contains(nextLink[1], `rel="next"`) {
				containsNext = true
				nextUrlRaw = nextLink[0]
				break
			}
		}

		if !containsNext {
			break
		}

		formattedUrl := strings.ReplaceAll(nextUrlRaw, "<", "")
		formattedUrl = strings.ReplaceAll(formattedUrl, ">", "")
		url = strings.Trim(formattedUrl, " ")

		body, err = io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		err = json.Unmarshal(body, &tempVariables)

		if err != nil {
			log.Fatalln(err)
		}

		variables = append(variables, tempVariables...)
	}

	jsonString, err := json.MarshalIndent(variables, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("variables.json")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(string(jsonString))

	if err2 != nil {
		log.Fatal(err2)
	}

	log.Println("Done")
}
