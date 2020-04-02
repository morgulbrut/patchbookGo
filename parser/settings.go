package parser

import (
	"strings"

	"github.com/morgulbrut/patchbookGo/patch"
)

func settings(s string) patch.Device {
	var d patch.Device
	s = strings.TrimPrefix(s, SETTINGS)
	s = strings.TrimSpace(s)
	if !(strings.HasSuffix(s, ":")) {
		ds := strings.Split(s, NAME_DEL)
		d.Name = strings.TrimSpace(ds[0])
		for _, st := range strings.Split(ds[1], SETTING) {
			d.Settings = append(d.Settings, setting(st))

		}
	}
	return d
}

func setting(s string) patch.Setting {
	var set patch.Setting
	pv := strings.Split(s, PARAM_DEL)
	set.Parameter = strings.TrimSpace(pv[0])
	set.Value = strings.TrimSpace(pv[1])
	return set
}
