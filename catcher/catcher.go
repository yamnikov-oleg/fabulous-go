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

type RepoInfo struct {
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

func MdHeaderItem(text string) (header string, lvl int, ok bool) {
	text = strings.TrimSpace(text)

	if len(text) < 3 {
		return "", 0, false
	}

	// Header
	if text[0] == '#' {
		return mdHeader(text)
	}

	// List item
	if isMdListItem(text) {
		return strings.TrimSpace(text[2:]), 8, true
	}

	// Bold text
	if isMdBoldLine(text) {
		return strings.TrimSpace(text[1 : len(text)-1]), 7, true
	}

	return "", 0, false
}

func isMdListItem(text string) bool {
	if len(text) < 3 {
		return false
	}

	fst, snd := text[0], text[1]
	if fst != '*' && fst != '+' && fst != '-' {
		return false
	}
	if snd != ' ' {
		return false
	}

	return true
}

func isMdBoldLine(text string) bool {
	if len(text) < 3 {
		return false
	}

	fst, snd, lst := text[0], text[1], text[len(text)-1]
	if fst != '*' || snd == ' ' || lst != '*' {
		return false
	}

	return true
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
	Titles   [8]string
}

type RepoList []*RepoEntry

func ReadRepoList(r io.Reader) (list RepoList, err error) {
	var titlesMap [8]string

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {

		line := strings.TrimSpace(scanner.Text())

		// There must always be at least 3 chars:
		// [# H], [+ H], [*H*]
		if len(line) < 3 {
			continue
		}

		if user, name, ok := MdRepoItem(line); ok && isMdListItem(line) {
			list = append(list, &RepoEntry{user, name, titlesMap})
			continue
		}

		if header, level, ok := MdHeaderItem(line); ok {
			titlesMap[level-1] = header
			for i := level; i < len(titlesMap); i++ {
				titlesMap[i] = ""
			}
			continue
		}

	} // for scanner.Scan()

	err = scanner.Err()
	return
}

func main() {

}
