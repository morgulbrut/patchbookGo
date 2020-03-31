package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/morgulbrut/patchbookGo/parser"
)

var f string // filename to open

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	flag.StringVar(&f, "f", "", "open file")
	flag.Parse()

	lines, err := readLines(f)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	patch := parser.File(lines)

	fmt.Println(patch)
}
