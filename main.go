package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"

	"github.com/morgulbrut/patchbookGo/parser"
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

	//fmt.Println(s)
	//fmt.Println("---------------------")
	re := regexp.MustCompile(`\n(\s*)\|`)
	dd := re.ReplaceAllString(s, " |")

	re2 := regexp.MustCompile(`\:(\s*)\|`)

	dd = re2.ReplaceAllString(dd, ":")
	//fmt.Println(dd)

	patch := parser.File(strings.Split(dd, "\n"))

	fmt.Println(patch)
}
