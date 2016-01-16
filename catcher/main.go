package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	AwesomeGoReadmeUrl = "https://raw.githubusercontent.com/avelino/awesome-go/master/REDME.md"
)

var (
	tmplFilename string
	outFilename  string
)

func init() {
	flag.StringVar(&tmplFilename, "t", "TEMPLATE.md", "template to use for catalog compilation")
	flag.StringVar(&outFilename, "o", "README.md", "output file")
}

func errduring(activity string, err interface{}, v ...interface{}) {
	t := fmt.Sprintf("Error during %v:\n%v\n", activity, err)
	fmt.Printf(t, v...)
	os.Exit(1)
}

func main() {
	flag.Parse()
}
