package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
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

func TestMarkdownHeaderParsing(t *testing.T) {
	mustParse := func(what string, level int, input string, expected string) {
		header, lvl, ok := MdHeader(input)
		expect(t, "Must parse "+what, ok, true)
		expect(t, "Must parse "+what, lvl, level)
		expect(t, "Must parse "+what, header, strings.TrimSpace(expected))
	}

	mustNotParse := func(what string, input string) {
		header, lvl, ok := MdHeader(input)
		expect(t, "Must not parse "+what, ok, false)
		expect(t, "Must not parse "+what, header, "")
		expect(t, "Must not parse "+what, lvl, 0)
	}

	mustNotParse("empty string", "")
	mustNotParse("empty header", "# ")
	mustNotParse("empty header with newline", "# \n")

	testHeaders := []string{"A header", "A n o t h e r header", "Nospaces", "spaces    "}
	for _, h := range testHeaders {
		mustNotParse("normal text", h)
		mustParse("normal header", 1, "# "+h, h)
		mustParse("normal header with a lot of spaces", 4, "####    "+h, h)
		mustNotParse("header with no space", "#"+h)
		mustParse("normal 3rd level header", 3, "### "+h, h)
		mustParse("normal 6th level header", 6, "###### "+h, h)
		mustNotParse("3rd level header with no space", "###"+h)
		mustParse("normal header with newline", 1, "# "+h+"\n", h)
		mustParse("normal 3rd level header with newline", 3, "### "+h+"\n", h)
		mustParse("normal 3rd level header with newline and spaces before newline", 3, "### "+h+"  \n", h)
		mustParse("normal 3rd level header with newline and spaces after newline", 3, "### "+h+"\n  ", h)
	}
}
