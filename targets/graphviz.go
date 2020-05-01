package targets

import (
	"regexp"
	"strings"

	"github.com/morgulbrut/patchbookGo/patch"
)

func Graphviz(p patch.Root) string {

	var sb strings.Builder
	sb.WriteString("digraph{\n")
	sb.WriteString("ordering=\"out\";\nrankdir=\"LR\";\nsplines=\"polyline\";\n")

	// Devices
	for _, d := range p.Devices {
		sb.WriteString(synthNode(d))
	}

	// Connections

	for _, c := range p.Conns {
		sb.WriteString(connection(c))
	}

	sb.WriteString("}")
	return sb.String()
}

func label(d patch.Device) string {
	var sb strings.Builder
	sb.WriteString("label=\"{")
	sb.WriteString(outputs(d.Outputs))
	sb.WriteString("|")
	sb.WriteString(settings(d.Settings, d.Name))
	sb.WriteString("|")
	sb.WriteString(inputs(d.Inputs))
	sb.WriteString("}")
	return sb.String()
}

func inputs(inps []string) string {
	var sb strings.Builder
	sb.WriteString("{")
	for i, inp := range inps {
		if i > 0 {
			sb.WriteString(" | ")
		}
		sb.WriteString(tag(inp))
	}
	sb.WriteString("}")
	return sb.String()
}

func outputs(outps []string) string {
	var sb strings.Builder
	sb.WriteString("{")
	for i, outp := range outps {
		if i > 0 {
			sb.WriteString(" | ")
		}
		sb.WriteString(tag(outp))
	}
	sb.WriteString("}")
	return sb.String()
}

func settings(sets []patch.Setting, name string) string {
	var sb strings.Builder
	sb.WriteString("{")
	sb.WriteString("{")
	sb.WriteString(name)
	sb.WriteString("}|{")
	for i, set := range sets {
		if i > 0 {
			sb.WriteString("\\n")
		}
		sb.WriteString(set.Parameter + " = " + set.Value)
	}
	sb.WriteString("}")
	sb.WriteString("}")
	return sb.String()
}

func synthNode(d patch.Device) string {
	// aethervco[label="{ {<_cv> CV}|{{AETHER VCO}|{Lfo Freq = 5\nLfo Pwm = 7}}| {<_pulse> PULSE | <_sub1> SUB 1 | <_sub2> SUB 2}}"  shape=Mrecord]
	var sb strings.Builder
	n := nodeName(d.Name)
	sb.WriteString(n + "[")
	sb.WriteString(label(d))
	sb.WriteString("\" shape=Mrecord]\n")
	return sb.String()
}

func connection(c patch.Connection) string {
	// aethervco:_pulse:e  ->  mixer:_ch1:w  [style=bold]

	linetypes := map[string]string{
		"CV":      "[color=gray]",
		"Audio":   "[style=bold]",
		"Pitch":   "[color=blue]",
		"Gate":    "[color=red, style=dashed]",
		"Trigger": "[color=orange, style=dashed]",
		"Clock":   "[color=purple, style=dashed]",
	}

	var sb strings.Builder
	sb.WriteString(nodeName(c.Source.Name))
	sb.WriteString(":")
	sb.WriteString(portName(c.Source.PortName))
	sb.WriteString(":e  -> ")
	sb.WriteString(nodeName(c.Dest.Name))
	sb.WriteString(":")
	sb.WriteString(portName(c.Dest.PortName))
	sb.WriteString(":w ")
	sb.WriteString(linetypes[c.Type])
	sb.WriteString("\n")
	return sb.String()
}

func sanString(s string) string {
	n := strings.ToLower(s)
	re := regexp.MustCompile(`[^a-z0-9]`)
	return re.ReplaceAllString(n, "")
}

func portName(s string) string {
	n := sanString(s)
	return "_" + n
}

func tag(s string) string {
	t := sanString(s)
	return "<_" + t + "> " + s
}

func nodeName(s string) string {
	n := sanString(s)
	return n
}
