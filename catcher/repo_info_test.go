package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func setupApiServer(requestedUrl *string, data interface{}) {
	f := func(w http.ResponseWriter, r *http.Request) {
		*requestedUrl = r.URL.String()
		dataEncoded, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}
		w.Write(dataEncoded)
	}
	server := httptest.NewServer(http.HandlerFunc(f))
	GithubApiUrl = server.URL
}

func setupErrorServer(code int, data interface{}) {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(code)
		if data == nil {
			return
		}
		dataEncoded, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}
		w.Write(dataEncoded)
	}
	server := httptest.NewServer(http.HandlerFunc(f))
	GithubApiUrl = server.URL
}

func TestRetrieveJson(t *testing.T) {
	setupErrorServer(http.StatusNotFound, nil)
	err := RetrieveJson(GithubApiUrl+"/url", new(int))
	assert(t, "Error must be non-nil", err != nil)

	setupErrorServer(http.StatusNotFound, 12)
	err = RetrieveJson(GithubApiUrl+"/url", new(int))
	assert(t, "Error must be non-nil", err != nil)
}

func TestRepoRequest(t *testing.T) {
	var (
		requestedUrl string
		repo         = &RepoInfo{}
	)
	repo.StargazersCount = 25565
	repo.Name = "somerepo"
	repo.Owner.Login = "someuser"

	setupApiServer(&requestedUrl, repo)

	repo2, err := RequestRepoInfo(repo.Owner.Login, repo.Name)

	assert(t,
		"There must be no error",
		err == nil,
	)

	expect(t,
		"Request URL must form correctly",
		requestedUrl,
		fmt.Sprintf("/%v/%v", repo.Owner.Login, repo.Name),
	)

	assertFatal(t,
		"Return value must not be nil",
		repo2 != nil,
	)

	expect(t,
		"Stargazers count must be retrieved correctly",
		repo2.StargazersCount,
		repo.StargazersCount,
	)

	expect(t,
		"Repo name must be retrieved correctly",
		repo2.Name,
		repo.Name,
	)

	expect(t,
		"Owner name must be retrieved correctly",
		repo2.Owner.Login,
		repo.Owner.Login,
	)
}

func TestCommitsRequest(t *testing.T) {
	var (
		requestedUrl string
		repo         = &RepoInfo{}
		comCount     = 20
		commits      []*CommitInfo
		page         = 45
	)

	repo.Name = "somerepo"
	repo.Owner.Login = "someuser"

	commits = make([]*CommitInfo, comCount)
	for i, _ := range commits {
		commits[i] = &CommitInfo{}
		commits[i].Sha = fmt.Sprintf("hailstalin%v", i)
		commits[i].Commit.Author.Date = time.Now().Add(time.Duration(-i) * time.Hour)
	}

	setupApiServer(&requestedUrl, commits)

	commits2, err := repo.RequestCommits(page)
	assertFatal(t, "Error must be nil", err == nil)

	expect(t,
		"Request url must be correct",
		requestedUrl,
		fmt.Sprintf("/%v/%v/commits?page=%v", repo.Owner.Login, repo.Name, page),
	)

	expectFatal(t, "Commits count be correct", len(commits2), comCount)
	for i, _ := range commits2 {
		expect(t,
			fmt.Sprintf("Commit #%v must be the same", i),
			*commits2[i], *commits[i],
		)
	}
}

func TestParticipationStatsRequest(t *testing.T) {
	var (
		requestedUrl string
		repo         = &RepoInfo{}
		stats        = &ParticipationStats{
			All:   []int{78, 8, 12, 51, 33, 23},
			Owner: []int{99, 123, 54, 234, 11, 0, 0, 0},
		}
	)
	repo.Name = "somerepo"
	repo.Owner.Login = "someuser"

	setupApiServer(&requestedUrl, stats)
	stats2, err := repo.RequestParticipationStats()

	assert(t, "Error must be nil", err == nil)
	expect(t,
		"Requested url must be correct",
		requestedUrl, fmt.Sprintf("/%v/%v/stats/participation", repo.Owner.Login, repo.Name),
	)
	assertFatal(t, "Stats must not be nil", stats2 != nil)

	expect(t, "Stats all commits count must be the same", len(stats2.All), len(stats.All))
	for i, _ := range stats2.All {
		expect(t,
			fmt.Sprintf("Stats commit of all #%d must be the same", i),
			stats2.All[i], stats.All[i])
	}

	expect(t, "Stats owner commits count must be the same", len(stats2.Owner), len(stats.Owner))
	for i, _ := range stats2.Owner {
		expect(t,
			fmt.Sprintf("Stats commit of all #%d must be the same", i),
			stats2.Owner[i], stats.Owner[i])
	}
}
