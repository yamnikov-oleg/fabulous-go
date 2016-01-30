package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTemplateDataComposition(t *testing.T) {
	entry := &RepoEntry{Username: "huares", Reponame: "mybolt"}

	infoUrl := fmt.Sprintf("/repos/%v/%v", entry.Username, entry.Reponame)
	infoRequested := false
	info := &RepoInfo{}
	info.Owner.Login = entry.Username
	info.Name = entry.Reponame
	info.StargazersCount = 909
	info.Description = "The little repo"

	statsUrl := infoUrl + "/stats/participation"
	statsRequested := false
	stats := &ParticipationStats{}
	stats.All = []int{5, 6, 7, 8, 10, 0, 0, 0, 12}
	commitsLastMonth := 5 + 6 + 7 + 8

	hf := func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.String()
		switch url {
		case infoUrl:
			infoRequested = true
			data, err := json.Marshal(info)
			if err != nil {
				panic(err)
			}
			w.Write(data)
		case statsUrl:
			statsRequested = true
			data, err := json.Marshal(stats)
			if err != nil {
				panic(err)
			}
			w.Write(data)
		default:
			w.WriteHeader(404)
		}
	}
	server := httptest.NewServer(http.HandlerFunc(hf))
	GithubApiUrl = server.URL

	err := ComposeTemplateData(entry)

	assert(t, "Error must be nil", err == nil)
	assertFatal(t, "Template data must not be nil", entry.TmplData != nil)

	expect(t, "Stars count must be correct", entry.TmplData.Stars, info.StargazersCount)
	expect(t, "Commits count must be correct", entry.TmplData.CommitsLastMonth, commitsLastMonth)
	expect(t, "Description must be correct", entry.TmplData.Description, info.Description)
}
