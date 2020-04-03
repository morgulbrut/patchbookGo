package targets

import (
	"fmt"
	"strings"

	"github.com/morgulbrut/patchbookGo/patch"

	"github.com/emicklei/dot"
)

func GraphvizTest() string {
	di := dot.NewGraph(dot.Directed)
	outside := di.Node("Outside")

	// A
	clusterA := di.Subgraph("Cluster A", dot.ClusterOption{})
	insideOne := clusterA.Node("one")
	insideTwo := clusterA.Node("two")

	// B
	clusterB := di.Subgraph("Cluster B", dot.ClusterOption{})
	insideThree := clusterB.Node("three")
	insideFour := clusterB.Node("four")

	// edges
	outside.Edge(insideFour).Edge(insideOne).Edge(insideTwo).Edge(insideThree).Edge(outside)

	return fmt.Sprintf(di.String())
}

func Graphviz(p patch.Root) string {
	g := dot.NewGraph(dot.Directed)
	g.Attr("rankdir", "LR")
	g.Attr("ordering", "out")
	g.Attr("splines", "polyline")

	// Devices
	for _, d := range p.Devices {
		//n := strings.ToLower(d.Name)
		//n = strings.ReplaceAll(n, " ", "")
		l := label(d)
		node := g.Node(l)
		node.Attr("shape", "Mrecord")

	}

	return fmt.Sprintf(g.String())
}

func label(d patch.Device) string {
	var sb strings.Builder
	sb.WriteString("{ {")
	for _, i := range d.Inputs {
		n := strings.ToLower(i)
		n = strings.ReplaceAll(n, " ", "")
		sb.WriteString(fmt.Sprintf("<_%s> %s | ", n, i))
	}
	sb.WriteString("} }")
	return sb.String()
}
