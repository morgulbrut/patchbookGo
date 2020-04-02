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
const SETTINGS = "*"
const SETTING = "|"
const NAME_DEL = ":"
const PARAM_DEL = "="

func File(s []string) patch.Root {
	var p patch.Root

	color256.PrintHiOrange("Parsing...:")
	for i, line := range s {
		fmt.Println(i, line)
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, CONN) || strings.HasPrefix(line, SETTINGS) {
			p.Devices = append(p.Devices, device(line))
		}

		if strings.HasPrefix(line, CONN) {
			p.Conns = append(p.Conns, connection(line))
		}

	}
	return p
}
