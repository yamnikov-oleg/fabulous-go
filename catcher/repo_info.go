package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	HttpClient = &http.Client{}
	// No trailing slash
	GithubApiUrl = "https://api.github.com"
)

type RepoInfo struct {
	Name  string `json:"name"`
	Owner struct {
		Login string `json:"login"`
	} `json:"owner"`
	StargazersCount int `json:"stargazers_count"`
}

func RequestRepoInfo(username string, reponame string) (*RepoInfo, error) {
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
	info := &RepoInfo{}
	err = json.Unmarshal(data, info)
	if err != nil {
		return nil, err
	}

	return info, nil
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
	// Form request
	url := fmt.Sprintf("%v/%v/%v/commits?page=%v", GithubApiUrl, r.Owner.Login, r.Name, page)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	// Run request
	resp, err := HttpClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// Buffer response
	buffer, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	// Unmarshal
	err = json.Unmarshal(buffer, &list)
	if err != nil {
		return
	}

	return
}

type ParticipationStats struct {
	All   []int `json:"all"`
	Owner []int `json:"owner"`
}

func (r *RepoInfo) RequestParticipationStats() (*ParticipationStats, error) {
	// Form a request
	url := fmt.Sprintf("%v/%v/%v/stats/participation", GithubApiUrl, r.Owner.Login, r.Name)
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
	stats := &ParticipationStats{}
	err = json.Unmarshal(data, stats)
	if err != nil {
		return nil, err
	}

	return stats, nil
}
