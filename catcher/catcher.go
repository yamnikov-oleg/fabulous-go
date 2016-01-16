package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	HttpClient = &http.Client{}
	// No trailing slash
	GithubApiUrl = "https://api.github.com"
)

type Repo struct {
	StargazersCount int `json:"stargazers_count"`
}

func ParseRepo(username string, reponame string) (*Repo, error) {
	// Form a request
	url := fmt.Sprintf("%v/%v/%v", GithubApiUrl, username, reponame)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/vnd.github.v3+json")

	// Perform the request
	resp, err := HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Buffer the response
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshal the response
	repo := &Repo{}
	err = json.Unmarshal(data, repo)
	if err != nil {
		return nil, err
	}

	return repo, nil
}

func MdHeader(text string) (header string, lvl int, ok bool) {
	if len(text) == 0 {
		return "", 0, false
	}
	if text[0] != '#' {
		return "", 0, false
	}

	spaceIndex := strings.Index(text, " ")
	if spaceIndex < 0 {
		return "", 0, false
	}

	marker := text[:spaceIndex]
	for _, rn := range marker {
		if rn != '#' {
			return "", 0, false
		}
	}

	lvl = len(marker)
	header = strings.TrimSpace(text[spaceIndex:])
	if header == "" {
		return "", 0, false
	}
	ok = true

	return
}

func main() {

}
