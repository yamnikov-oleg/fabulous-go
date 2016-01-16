package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var (
	HttpClient = &http.Client{}
	// No trailing slash
	GithubApiUrl = "https://api.github.com"
)

func RetrieveJson(url string, outData interface{}) error {
	// Form a request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Accept", "application/vnd.github.v3+json")

	// Perform the request
	resp, err := HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status code not OK (%v)", resp.Status)
	}

	// Unmarshal the response
	err = json.NewDecoder(resp.Body).Decode(outData)
	if err != nil {
		return err
	}

	return nil
}

type RepoInfo struct {
	Name  string `json:"name"`
	Owner struct {
		Login string `json:"login"`
	} `json:"owner"`
	StargazersCount int `json:"stargazers_count"`
}

func RequestRepoInfo(username string, reponame string) (info *RepoInfo, err error) {
	url := fmt.Sprintf("%v/%v/%v", GithubApiUrl, username, reponame)
	info = &RepoInfo{}
	err = RetrieveJson(url, info)
	return
}

type CommitInfo struct {
	Sha    string `json:"sha"`
	Commit struct {
		Author struct {
			Date time.Time `json:"date"`
		} `json:"author"`
	} `json:"commit"`
}

func (r *RepoInfo) RequestCommits(page int) (list []*CommitInfo, err error) {
	url := fmt.Sprintf("%v/%v/%v/commits?page=%v", GithubApiUrl, r.Owner.Login, r.Name, page)
	err = RetrieveJson(url, &list)
	return
}

type ParticipationStats struct {
	All   []int `json:"all"`
	Owner []int `json:"owner"`
}

func (r *RepoInfo) RequestParticipationStats() (stats *ParticipationStats, err error) {
	url := fmt.Sprintf("%v/%v/%v/stats/participation", GithubApiUrl, r.Owner.Login, r.Name)
	stats = &ParticipationStats{}
	err = RetrieveJson(url, stats)
	return
}
