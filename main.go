package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
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

	fmt.Println(s)
	fmt.Println("---------------------")
	re := regexp.MustCompile(`\n(\s*)\|`)
	dd := re.ReplaceAllString(s, " |")

	fmt.Println(dd)
	/*
		lines, err := helferlein.ReadLines(f)
		if err != nil {
			log.Fatalf("readLines: %s", err)
		}
	*/

	//patch := parser.File(lines)

	//fmt.Println(patch)
}
