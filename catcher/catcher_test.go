package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
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

func expectFatal(t *testing.T, logText string, actual interface{}, expected interface{}) {
	if actual != expected {
		t.Fatalf("%v. Expected %#v, got %#v.", logText, expected, actual)
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
		repo         = &RepoInfo{
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
	repo2, err := RequestRepoInfo(username, reponame)

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
		header, lvl, ok := MdHeaderItem(input)
		expect(t, "Must parse "+what, ok, true)
		expect(t, "Must parse "+what, lvl, level)
		expect(t, "Must parse "+what, header, strings.TrimSpace(expected))
	}

	mustNotParse := func(what string, input string) {
		header, lvl, ok := MdHeaderItem(input)
		expect(t, "Must not parse "+what, ok, false)
		expect(t, "Must not parse "+what, header, "")
		expect(t, "Must not parse "+what, lvl, 0)
	}

	mustNotParse("empty string", "")
	mustNotParse("empty header", "# ")
	mustNotParse("empty header with newline", "# \n")

	testHeaders := []string{"A header", "Nospaces", "spaces    "}
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

		mustParse("bold text as 7th level header", 7, "*"+h+"*", h)
		mustNotParse("bold text with no leading asterisk", "*"+h)
		mustParse("astreisk list item as 8th level header", 8, "* "+h+"*", h+"*")
		mustParse("hyphen list item as 8th level header", 8, "- "+h, h)
		mustParse("plus list item as 8th level header", 8, "+ "+h, h)
	}
}

func TestRepoItemParsing(t *testing.T) {
	mustParse := func(what string, input string, username string, reponame string) {
		un, rn, ok := MdRepoItem(input)
		expect(t, "Must parse "+what, ok, true)
		expect(t, "Must parse "+what, un, username)
		expect(t, "Must parse "+what, rn, reponame)
	}

	mustNotParse := func(what string, input string) {
		un, rn, ok := MdRepoItem(input)
		expect(t, "Must not parse "+what, ok, false)
		expect(t, "Must not parse "+what, un, "")
		expect(t, "Must not parse "+what, rn, "")
	}

	mustNotParse("empty string", "")
	mustNotParse("string with no repo link", "   Some string   ")
	mustParse("string with repo link", "(http://github.com/moi/monrepo)", "moi", "monrepo")
	mustNotParse("string with wrong link", "(http://github.org/moi/monrepo)")
	mustParse("string with repo link and label", "[Some label](http://github.com/moi/monrepo)", "moi", "monrepo")
	mustParse("string with https repo link", "(https://github.com/moi/monrepo)", "moi", "monrepo")
	mustParse("string with repo link and text", "Text before (https://github.com/moi/monrepo) text after", "moi", "monrepo")
	mustParse("string with repo link, label and text", "* [a label](https://github.com/moi/monrepo) text after", "moi", "monrepo")
	mustParse("string with several links", "* [wrong link](https://example.com/path/) text after [a label](https://github.com/moi/monrepo)", "moi", "monrepo")
	mustParse("first repo if several", "* [wrong link](https://github.com/moi/monrepo) text after [a label](https://github.com/toi/tonrepo)", "moi", "monrepo")
	mustNotParse("no repo github link (less path segments)", "(http://github.com/moi)")
	mustNotParse("no repo github link (more path segments)", "(http://github.com/moi/monrepo/extra)")
}

func tryOpen(pathlist ...string) (f *os.File, err error) {
	if len(pathlist) == 0 {
		return nil, errors.New("path list empty")
	}

	for _, path := range pathlist {
		f, err = os.Open(path)
		if err == nil {
			return
		}
	}
	return
}

func TestRepoListParsing(t *testing.T) {
	// Try opening in several paths, because wording dir may change
	file, err := tryOpen("repo_list_test_sample.md", "catcher/repo_list_test_sample.md")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	list, err := ReadRepoList(file)

	assert(t,
		"Error must be nil",
		err == nil,
	)
	expectFatal(t,
		"Repo list must have correct length",
		len(list),
		34,
	)

	assertEntry := func(what string, e *RepoEntry, user string, name string, titles ...string) {
		t.Logf("Testing %v: %v", what, e)
		assert(t,
			fmt.Sprintf("%v must be %v/%v", what, user, name),
			e.Username == user && e.Reponame == name,
		)
		assert(t,
			fmt.Sprintf("%v must have %v titles", what, len(titles)),
			len(e.Titles) == len(titles),
		)
		if len(e.Titles) != len(titles) {
			return
		}
		for i, _ := range titles {
			assert(t,
				fmt.Sprintf("%v %dth title must be %v", what, i, titles[i]),
				e.Titles[i] == titles[i],
			)
		}
	}

	assertEntry(
		"first entry", list[0], "user", "package1",
		"Site Header",
		"Simple category",
		"", "", "", "",
		"Sample description",
		"",
	)

	assertEntry(
		"entry under subheader", list[4], "user", "package2",
		"Site Header",
		"Category with subheaders",
		"Subheader 1",
		"", "", "",
		"Description",
		"",
	)

	assertEntry(
		"entry under third category", list[11], "user", "package3",
		"Site Header",
		"Simple category after complex one",
		"", "", "", "",
		"Description",
		"",
	)

	assertEntry(
		"entry under section", list[15], "user", "package1",
		"Site Header",
		"Category with two sections",
		"", "", "", "",
		"Section 2 as description",
		"",
	)

	assertEntry(
		"entry under list section", list[19], "user", "package2",
		"Site Header",
		"Category with list sections",
		"", "", "", "",
		"Common description",
		"Section 1 as list item",
	)

	assertEntry(
		"entry under second simple cat", list[26], "user", "package3",
		"Site Header",
		"Simple category after complex one",
		"", "", "", "",
		"Description",
		"",
	)

	assertEntry(
		"entry with broken indentation", list[30], "user", "package4",
		"Site Header",
		"Category with broken indentation",
		"", "", "", "",
		"Description",
		"",
	)

	assertEntry(
		"entry after unparsed one", list[31], "user", "package1",
		"Another FIRST LEVEL header",
		"Category",
		"Subcategory",
		"", "", "", "", "",
	)

}
