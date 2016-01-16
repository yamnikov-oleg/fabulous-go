package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

var (
	HttpClient = &http.Client{}
	// No trailing slash
	GithubApiUrl     = "https://api.github.com"
	MdRepoLinkRegexp = regexp.MustCompile(`\(https?:\/\/github.com\/(\w+)\/(\w+)\)`)
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

func MdHeaderItem(text string) (header string, lvl int, ok bool) {
	text = strings.TrimSpace(text)

	// There must always be at least 3 chars:
	// [# H], [+ H], [*H*]
	if len(text) < 3 {
		return "", 0, false
	}

	fst, snd, lst := text[0], text[1], text[len(text)-1]

	// Header
	if fst == '#' {
		return mdHeader(text)
	}

	// List item
	if (fst == '*' || fst == '+' || fst == '-') && snd == ' ' {
		return strings.TrimSpace(text[2:]), 8, true
	}

	// Bold text
	if fst == '*' && snd != ' ' && lst == '*' {
		return strings.TrimSpace(text[1 : len(text)-1]), 7, true
	}

	return "", 0, false
}

func mdHeader(text string) (header string, lvl int, ok bool) {
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

func MdRepoItem(text string) (username string, reponame string, ok bool) {
	match := MdRepoLinkRegexp.FindStringSubmatch(text)
	if len(match) == 0 {
		return "", "", false
	}
	return match[1], match[2], true
}

type RepoEntry struct {
	Username string
	Reponame string
	Titles   []string
}

type RepoList []*RepoEntry

func ReadRepoList(r io.Reader) (list RepoList, err error) {
	var titlesMap [8]string

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) < 3 {
			continue
		}

		user, name, ok := MdRepoItem(line)

		fst, snd := line[0], line[1]
		if (fst == '*' || fst == '+' || fst == '-') && snd == ' ' && ok {

			titles := make([]string, 0, len(titlesMap))
			for _, t := range titlesMap {
				if t != "" {
					titles = append(titles, t)
				}
			}
			list = append(list, &RepoEntry{user, name, titles})
			continue
		}

		header, level, ok := MdHeaderItem(line)
		if ok {
			titlesMap[level-1] = header
			for i := level; i < len(titlesMap); i++ {
				titlesMap[i] = ""
			}
			continue
		}

	}

	err = scanner.Err()
	return
}

func main() {

}
