package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"
)

const (
	AwesomeGoReadmeUrl = "https://raw.githubusercontent.com/avelino/awesome-go/master/README.md"
)

var (
	outFilename string
)

func init() {
	flag.StringVar(&outFilename, "o", "repos.json", "output file")
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

	fmt.Printf("Performing `GET %v`...\n", AwesomeGoReadmeUrl)
	resp, err := http.Get(AwesomeGoReadmeUrl)
	if err != nil {
		fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fatalf("status code not successfull (%v)", resp.Status)
	}

	fmt.Printf("Parsing repo catalog from the response body...\n")
	catalog, err := ReadRepoCatalog(resp.Body)
	if err != nil {
		fatal(err)
	}

	fmt.Printf("Parsed %v blocks of repo entries\n", len(catalog))
	if len(catalog) == 0 {
		fatalf("no entries parsed! Must be an error.")
	}

	for i, block := range catalog {
		fmt.Printf("Composing data for block [%v/%v]", i+1, len(catalog))
		for _, entry := range block.Entries {
			fmt.Print(".")
			err := ComposeTemplateData(entry)
			if err != nil {
				fmt.Println()
				fatal(err)
			}
			// Rate limit
			time.Sleep(2 * time.Second)
		}
		fmt.Println()
	}

	fmt.Printf("Creating %v...\n", outFilename)
	out, err := os.Create(outFilename)
	if err != nil {
		fatal(err)
	}
	defer out.Close()

	fmt.Printf("Encoding data...\n")
	err = json.NewEncoder(out).Encode(catalog)
	if err != nil {
		fatal(err)
	}

	fmt.Println("COMPLETE")
}
