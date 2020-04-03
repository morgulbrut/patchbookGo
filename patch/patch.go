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
	sb.WriteString("\nDevices:\n")
	for i, d := range p.Devices {
		sb.WriteString(fmt.Sprintf("%d\t%s:\n", i+1, d.Name))
		if len(d.Settings) > 0 {
			sb.WriteString("\t\tSettings:\n")
			for i, s := range d.Settings {
				sb.WriteString(fmt.Sprintf("\t\t%d\t%s: %s\n", i+1, s.Parameter, s.Value))
			}
		}
		if len(d.Inputs) > 0 {
			sb.WriteString("\t\tInputs:\n")
			for i, s := range d.Inputs {
				sb.WriteString(fmt.Sprintf("\t\t%d\t%s\n", i+1, s))
			}
		}
		if len(d.Outputs) > 0 {
			sb.WriteString("\t\tOutputs:\n")
			for i, s := range d.Outputs {
				sb.WriteString(fmt.Sprintf("\t\t%d\t%s\n", i+1, s))
			}
		}
	}

	sb.WriteString("\nConnections:\n")
	for i, c := range p.Conns {
		sb.WriteString(fmt.Sprintf("%d\t%s (%s) -> %s (%s), T: %s\n", i+1, c.Source.Name, c.Source.PortName, c.Dest.Name, c.Dest.PortName, c.Type))
	}
	return sb.String()
}
