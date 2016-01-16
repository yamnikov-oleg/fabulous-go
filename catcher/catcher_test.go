package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupApiServer(f http.HandlerFunc) {
	server := httptest.NewServer(f)
	GithubApiUrl = server.URL
}

func expect(t *testing.T, logText string, actual interface{}, expected interface{}) {
	if actual != expected {
		t.Errorf("%v. Expected %#v, got %#v.", logText, expected, actual)
	}
}

func assert(t *testing.T, logText string, cond bool) {
	if !cond {
		t.Error(logText)
	}
}

func assertFatal(t *testing.T, logText string, cond bool) {
	if !cond {
		t.Fatal(logText)
	}
}

func TestRepoParsing(t *testing.T) {
	var (
		requestedUrl string
		repo         = &Repo{
			StargazersCount: 25565,
		}
	)
	setupApiServer(func(w http.ResponseWriter, r *http.Request) {
		requestedUrl = r.URL.String()
		repoEncoded, err := json.Marshal(repo)
		if err != nil {
			panic(err)
		}
		w.Write(repoEncoded)
	})

	username, reponame := "someuser", "somerepo"
	repo2, err := ParseRepo(username, reponame)

	assert(t,
		"There must be no error",
		err == nil,
	)

	expect(t,
		"Request URL must form correctly",
		requestedUrl,
		fmt.Sprintf("/%v/%v", username, reponame),
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
}
