package internals

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func FetchVariables(host string, projectId string, privateToken string) ([]Variable, error) {
	client := http.Client{}
	url := fmt.Sprintf("%s/api/v4/projects/%s/variables", host, projectId)
	var variables []Variable = []Variable{}

	for {
		var body []byte
		var tempVariables []Variable
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}

		req.Header.Add("PRIVATE-TOKEN", privateToken)

		resp, err := client.Do(req)
		if err != nil {
			return nil, err
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
			return nil, err
		}
		err = json.Unmarshal(body, &tempVariables)

		if err != nil {
			return nil, err
		}

		variables = append(variables, tempVariables...)
	}

	return variables, nil
}
