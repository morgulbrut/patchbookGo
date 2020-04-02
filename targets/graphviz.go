package targets

import (
	"fmt"

	"github.com/morgulbrut/patchbookGo/patch"

	"github.com/emicklei/dot"
)

type Graph struct {
	Devices []Device
	Conns   []patch.Connection
}

type Device struct {
	Name string
	Sets []patch.Setting
}

func Test() {
	g := dot.NewGraph(dot.Directed)
	n1 := g.Node("coding")
	n2 := g.Node("testing a little").Box()

	g.Edge(n1, n2)
	g.Edge(n2, n1, "back").Attr("color", "red")

	fmt.Println(g.String())
}
