package parser

import (
	"strings"

	"github.com/morgulbrut/patchbookGo/patch"
)

func settings(s string) patch.Settings {
	var set patch.Settings
	s = strings.TrimPrefix(s, SETTINGS)
	s = strings.TrimSpace(s)
	if !(strings.HasSuffix(s, ":")) {
		ds := strings.Split(s, NAME_DEL)
		set.Device = strings.TrimSpace(ds[0])
		for _, st := range strings.Split(ds[1], SETTING) {
			set.Sets = append(set.Sets, setting(st))

		}
	}
	return set
}

func setting(s string) patch.Setting {
	var set patch.Setting
	pv := strings.Split(s, PARAM_DEL)
	set.Parameter = strings.TrimSpace(pv[0])
	set.Value = strings.TrimSpace(pv[1])
	return set
}
