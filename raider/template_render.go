package main

type TmplRepoData struct {
	Stars            int
	CommitsLastMonth int
	Description      string
}

func ComposeTemplateData(e *RepoEntry) error {
	data := &TmplRepoData{}

	info, err := RequestRepoInfo(e.Username, e.Reponame)
	if err != nil {
		return err
	}
	data.Stars = info.StargazersCount
	data.Description = info.Description

	stat, err := info.RequestParticipationStats()
	if err != nil {
		return err
	}
	for _, c := range stat.All[:4] {
		data.CommitsLastMonth += c
	}

	e.TmplData = data
	return nil
}
