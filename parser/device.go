package parser

import (
	"strings"

	"github.com/morgulbrut/patchbookGo/patch"
)

func deviceSettings(s string) patch.Device {
	var d patch.Device
	s = strings.Split(s, NAME_DEL)[0]
	s = strings.Trim(s, SETTINGS)
	s = strings.TrimSpace(s)
	d.Name = s
	return d
}

func deviceInput(s string) patch.Device {
	var d patch.Device
	s = strings.Split(s, "(")[0]
	s = strings.Trim(s, CONN)
	s = strings.TrimSpace(s)
	d.Name = s
	return d
}

func deviceOutput(s string) patch.Device {
	var d patch.Device
	out := strings.Split(s, ">")
	s = out[len(out)-1]
	s = strings.Split(s, "(")[0]
	s = strings.TrimSpace(s)
	d.Name = s
	return d
}
