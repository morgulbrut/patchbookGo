package parser

import (
	"strings"

	"github.com/morgulbrut/patchbookGo/patch"
)

func device(s string) patch.Device {
	var d patch.Device
	if strings.HasPrefix(s, CONN) {
		s = strings.Split(s, "(")[0]
		s = strings.Trim(s, CONN)
	} else {
		s = strings.Split(s, NAME_DEL)[0]
		s = strings.Trim(s, SETTINGS)
	}
	//	s = strings.ReplaceAll(s, " ", "")
	//	s = strings.ToLower(s)
	s = strings.TrimSpace(s)
	d.Name = s

	return d
}
