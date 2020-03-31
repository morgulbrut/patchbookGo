package parser

import (
	"fmt"
	"strings"

	"github.com/morgulbrut/color256"
)

func File(lines []string) Patch {
	color256.PrintHiOrange("Parsing...:")
	for i, line := range lines {
		fmt.Println(i, line)
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line,"*")
	}

	return Patch{}
}
