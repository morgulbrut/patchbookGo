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

		// Add new device if not allready there
		if strings.HasPrefix(line, CONN) || strings.HasPrefix(line, SETTINGS) {
			dev := device(line)
			found := false
			for _, d := range p.Devices {
				if d.Name == dev.Name {
					found = true
				}
			}
			if !found {
				p.Devices = append(p.Devices, dev)
			}
		}

		// Add setting to devices
		if strings.HasPrefix(line, SETTINGS) {
			dev := device(line)
			for _, d := range p.Devices {
				if d.Name == dev.Name {
					sets := settings(line)
					color256.PrintBgGreen(sets)
					d.Settings = append(d.Settings, sets...)
				}
			}
		}

		if strings.HasPrefix(line, CONN) {
			p.Conns = append(p.Conns, connection(line))
		}

	}
	return p
}
