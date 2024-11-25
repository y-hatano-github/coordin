package main

import (
	"fmt"

	c "github.com/y-hatano-github/coordin"
)

func main() {
	var m [14][14]string
	for i, rows := range m {
		for j := range rows {
			m[i][j] = "."
		}
	}

	ps := c.Line(c.Point{X: 2, Y: 2}, c.Point{X: 11, Y: 11})
	for _, p := range ps {
		m[p.X][p.Y] = "*"
	}
	for i, rows := range m {
		for j := range rows {
			fmt.Print(m[i][j])
		}
		fmt.Println()
	}
}
