package patch

import (
	"fmt"
	"strings"
)

type Patch struct {
	Conns []Connection
	Sets  []Settings
}

func (p Patch) String() string {
	var sb strings.Builder

	sb.WriteString("\n------------------------\nPatch\n------------------------\n")
	sb.WriteString("\nConnections:\n")
	for i, c := range p.Conns {
		sb.WriteString(fmt.Sprintf("%d\t%s (%s) -> %s (%s), T: %s\n", i+1, c.Source.Device, c.Source.PortName, c.Dest.Device, c.Dest.PortName, c.Type))
	}
	sb.WriteString("\nSettings:\n")
	for i, s := range p.Sets {
		sb.WriteString(fmt.Sprintf("%d\t%s:\n", i+1, s.Device))
		for j, t := range s.Sets {
			sb.WriteString(fmt.Sprintf("\t%d\t%s: %s\n", j+1, t.Parameter, t.Value))
		}
	}
	return sb.String()
}

type Connection struct {
	Source Port
	Dest   Port
	Type   string
}

type Port struct {
	Device   string
	PortName string
}

type Settings struct {
	Device string
	Sets   []Setting
}

type Setting struct {
	Parameter string
	Value     string
}
