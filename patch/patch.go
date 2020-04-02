package patch

import (
	"fmt"
	"strings"
)

type Root struct {
	Conns   []Connection
	Devices []Device
}

type Connection struct {
	Source Port
	Dest   Port
	Type   string
}

type Port struct {
	Name     string
	PortName string
}

type Device struct {
	Name     string
	Inputs   []string
	Outputs  []string
	Settings []Setting
}

type Setting struct {
	Parameter string
	Value     string
}

// Stringer: returns a patch.Root object in a formated way
func (p Root) String() string {
	var sb strings.Builder

	sb.WriteString("\n------------------------\nPatch\n------------------------\n")
	sb.WriteString("\nConnections:\n")
	for i, c := range p.Conns {
		sb.WriteString(fmt.Sprintf("%d\t%s (%s) -> %s (%s), T: %s\n", i+1, c.Source.Name, c.Source.PortName, c.Dest.Name, c.Dest.PortName, c.Type))
	}
	sb.WriteString("\nSettings:\n")
	for i, s := range p.Devices {
		sb.WriteString(fmt.Sprintf("%d\t%s:\n", i+1, s.Name))
		for j, t := range s.Settings {
			sb.WriteString(fmt.Sprintf("\t%d\t%s: %s\n", j+1, t.Parameter, t.Value))
		}
	}
	return sb.String()
}
