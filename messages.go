package main

import (
	"fmt"

	"github.com/gookit/color"
)

type Tip struct {
	Text string
}

func Comment(comment string) string {
	return color.FgDarkGray.Text(fmt.Sprintf("// %s", comment))
}
