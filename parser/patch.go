package parser

type Patch struct {
	Conns []Connection
	Sets  []Settings
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
