package main

import (
	"flag"
	"io/ioutil"
	"log"
	"regexp"
	"strings"

	"github.com/morgulbrut/patchbookGo/parser"
	"github.com/morgulbrut/patchbookGo/targets"
)

var f string // filename to open

func main() {
	flag.StringVar(&f, "f", "", "open file")
	flag.Parse()

	buf, err := ioutil.ReadFile(f)
	if err != nil {
		log.Fatal(err)
	}
	s := string(buf)

	re := regexp.MustCompile(`\n(\s*)\|`)
	s = re.ReplaceAllString(s, " |")

	re2 := regexp.MustCompile(`\:(\s*)\|`)
	s = re2.ReplaceAllString(s, ":")

	patch := parser.File(strings.Split(s, "\n"))

	// fmt.Println(patch)

	dot := targets.Graphviz(patch)
	// fmt.Println(targets.Graphviz(patch))

	err = ioutil.WriteFile("patch.txt", []byte(dot), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
