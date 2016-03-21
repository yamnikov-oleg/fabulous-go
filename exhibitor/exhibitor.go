package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
	"text/template"
)

type TmplRepoData struct {
	Stars            int
	CommitsLastMonth int
	Description      string
}

type RepoEntry struct {
	Username string
	Reponame string
	TmplData TmplRepoData
}

func (e *RepoEntry) Name() string {
	return e.Username + "/" + e.Reponame
}

type RepoEntryBlock struct {
	Entries []RepoEntry
	Titles  [8]string
}

type RepoCatalog []RepoEntryBlock

var (
	jsonFilename string
	tmplFilename string
	outFilename  string
)

func init() {
	flag.StringVar(&jsonFilename, "j", "repos.json", "File containing json data to pass to template")
	flag.StringVar(&tmplFilename, "t", "TEMPLATE.md", "Template file")
	flag.StringVar(&outFilename, "o", "README.md", "Output file")
}

type ByName []RepoEntry

func (l ByName) Len() int           { return len(l) }
func (l ByName) Less(i, j int) bool { return l[i].Name() < l[j].Name() }
func (l ByName) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }

type ByReponame []RepoEntry

func (l ByReponame) Len() int           { return len(l) }
func (l ByReponame) Less(i, j int) bool { return l[i].Reponame < l[j].Reponame }
func (l ByReponame) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }

type ByStarsDesc []RepoEntry

func (l ByStarsDesc) Len() int           { return len(l) }
func (l ByStarsDesc) Less(i, j int) bool { return l[i].TmplData.Stars > l[j].TmplData.Stars }
func (l ByStarsDesc) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }

type SortModeFunc func([]RepoEntry) sort.Interface

var SortModes = map[string]SortModeFunc{
	"name":     func(l []RepoEntry) sort.Interface { return ByName(l) },
	"reponame": func(l []RepoEntry) sort.Interface { return ByReponame(l) },
	"stars":    func(l []RepoEntry) sort.Interface { return ByStarsDesc(l) },
}

var Smode string

func init() {
	flag.StringVar(&Smode, "sort", "name", "Sort repositories by `field` inside each repo block.\n\tPossible values: 'name', 'reponame', 'stars'.")
}

func fatal(err interface{}) {
	fmt.Printf("[ERROR] %v\n", err)
	os.Exit(1)
}

func fatalf(format string, v ...interface{}) {
	t := fmt.Sprintf(format, v...)
	fmt.Printf("[ERROR] %v\n", t)
	os.Exit(1)
}

func main() {
	flag.Parse()

	fmt.Println("Compiling template...")
	funcmap := template.FuncMap{
		"title": func(i int, t string) string {
			if t == "" {
				return ""
			}

			if i < 6 {
				return strings.Repeat("#", i+1) + " " + t + "\n"
			} else if i == 6 {
				return "*" + t + "*\n"
			} else if i == 7 {
				return "* " + t + "\n"
			}

			return ""
		},
		"header_link": func(t string) string {
			buf := []byte(t)
			buf = bytes.ToLower(buf)

			link := make([]byte, 0, len(buf))
			link = append(link, '#')

			for _, b := range buf {
				if b >= 'a' && b <= 'z' {
					link = append(link, b)
					continue
				}
				if b == ' ' || b == '-' {
					link = append(link, '-')
					continue
				}
			}

			return string(link)
		},
		"times": func(count int, s string) string {
			buf := make([]byte, 0, len(s)*count)
			b := []byte(s)
			for i := 0; i < count; i++ {
				buf = append(buf, b...)
			}
			return string(buf)
		},
	}

	tmpl, err := template.New("").Funcs(funcmap).ParseFiles(tmplFilename)
	if err != nil {
		fatal(err)
	}

	fmt.Println("Parsing json file...")
	jsonFile, err := os.Open(jsonFilename)
	if err != nil {
		fatal(err)
	}

	var data RepoCatalog
	err = json.NewDecoder(jsonFile).Decode(&data)
	if err != nil {
		fatal(err)
	}
	if len(data) == 0 {
		fatal("No data was read. There must be an error")
	}
	data[0].Titles[0] = "Packages"

	sortFunc, ok := SortModes[Smode]
	if !ok {
		fatalf("No sort mode %q defined", Smode)
	}
	for _, b := range data {
		sort.Sort(sortFunc(b.Entries))
	}

	fmt.Println("Rendering template into file...")
	outFile, err := os.Create(outFilename)
	if err != nil {
		fatal(err)
	}

	err = tmpl.ExecuteTemplate(outFile, tmplFilename, data)
	if err != nil {
		fatal(err)
	}

	fmt.Println("COMPLETE")
}
