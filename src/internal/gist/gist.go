package gist

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const Url = "https://api.github.com/gists"

func Create(token, filename, description, content string) (string, error) {
	if description == "" {
		description = "Improve yourself using this Gist"
	}

	body := CreateRequest{
		Description: description,
		Public:      true,
		Files: map[string]File{
			filename: {Content: content},
		},
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, Url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	addHeaders(req, token)

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var resp CreateResponse

	err = json.Unmarshal(responseBody, &resp)
	if err != nil {
		return "", err
	}

	return strings.ReplaceAll(resp.URL, "\n", ""), nil
}

func addHeaders(r *http.Request, token string) {
	r.Header.Set("Accept", "application/vnd.github+json")
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	r.Header.Set("X-GitHub-Api-Version", "2022-11-28")
}
