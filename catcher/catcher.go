package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	HttpClient = &http.Client{}
	ServerUrl  = "https://api.github.com"
)

type Repo struct {
	StargazersCount int `json:"stargazers_count"`
}

func ParseRepo(username string, reponame string) (*Repo, error) {
	url := fmt.Sprintf("%v/%v/%v", ServerUrl, username, reponame)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/vnd.github.v3+json")

	resp, err := HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	repo := &Repo{}
	err = json.Unmarshal(data, repo)
	if err != nil {
		return nil, err
	}

	return repo, nil
}

func main() {

}
