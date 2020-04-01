package parser

import (
	"strings"

	"github.com/morgulbrut/patchbookGo/patch"
)

func connection(s string) patch.Connection {
	var c patch.Connection
	var sd []string
	s = strings.TrimPrefix(s, CONN)
	if strings.Contains(s, CONN_CV) {
		sd = strings.Split(s, CONN_CV)
		c.Type = "CV"
	}
	if strings.Contains(s, CONN_AUDIO) {
		sd = strings.Split(s, CONN_AUDIO)
		c.Type = "Audio"
	}
	if strings.Contains(s, CONN_GATE) {
		sd = strings.Split(s, CONN_GATE)
		c.Type = "Gate"
	}
	if strings.Contains(s, CONN_PITCH) {
		sd = strings.Split(s, CONN_PITCH)
		c.Type = "Pitch"
	}
	if strings.Contains(s, CONN_TRIGGER) {
		sd = strings.Split(s, CONN_TRIGGER)
		c.Type = "Trigger"
	}
	if strings.Contains(s, CONN_CLOCK) {
		sd = strings.Split(s, CONN_CLOCK)
		c.Type = "Clock"
	}
	c.Source = port(sd[0])
	c.Dest = port(sd[1])
	return c
}

func port(s string) patch.Port {
	var p patch.Port
	np := strings.Split(s, "(")
	p.Device = strings.Trim(np[0], " )")
	p.PortName = strings.Trim(np[1], " )")
	return p
}
