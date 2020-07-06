package utils

import (
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

// ReadFile reads and sanitzes a patchbook file
func ReadFile(fn string) []string {
	buf, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Fatal(err)
	}
	s := string(buf)

	re := regexp.MustCompile(`\n(\s*)\|`)
	s = re.ReplaceAllString(s, " |")

	re2 := regexp.MustCompile(`\:(\s*)\|`)
	s = re2.ReplaceAllString(s, ":")

	return strings.Split(s, "\n")
}
