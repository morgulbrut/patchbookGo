package patch

import "fmt"

type Patch struct {
	Conns []Connection
	Sets  []Settings
}

func (p Patch) String() string {
	return fmt.Sprintf("Connections: \n%v \nSettings: \n%v", p.Conns, p.Sets)
}

type Connection struct {
	Source Port
	Dest   Port
	Type   string
}

func (c Connection) String() string {
	return fmt.Sprintf("\tSource: %v, Destination: %v, Type: %s \n", c.Source, c.Dest, c.Type)
}

type Port struct {
	Device   string
	PortName string
}

func (p Port) String() string {
	return fmt.Sprintf("%s (%s)", p.Device, p.PortName)
}

type Settings struct {
	Device string
	Sets   []Setting
}

func (s Settings) String() string {
	return fmt.Sprintf("\n\tDevice: %v \n\t\tSettings: \n%v", s.Device, s.Sets)
}

type Setting struct {
	Parameter string
	Value     string
}

func (s Setting) String() string {
	return fmt.Sprintf("\t\t%s = %s\n", s.Parameter, s.Value)
}
