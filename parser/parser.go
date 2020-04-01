package parser

import (
	"fmt"
	"strings"

	"github.com/morgulbrut/color256"
	"github.com/morgulbrut/patchbookGo/patch"
)

const CONN = "-"
const CONN_CV = ">>"
const CONN_AUDIO = "->"
const CONN_PITCH = "p>" //1v/oct or Hz/V
const CONN_GATE = "g>"
const CONN_TRIGGER = "g>"
const CONN_CLOCK = "g>"

func File(s []string) patch.Patch {
	var p patch.Patch

	color256.PrintHiOrange("Parsing...:")
	for i, line := range s {
		fmt.Println(i, line)
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, CONN) {
			p.Conns = append(p.Conns, connection(line))
		}
	}
	return p
}

func connection(s string) patch.Connection {
	var c patch.Connection
	if strings.Contains(s, CONN_CV) {
		c.Type = "CV"
	}
	return c
}
