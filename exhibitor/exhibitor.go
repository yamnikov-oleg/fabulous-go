package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
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
