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
