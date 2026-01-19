package main

import (
	"fmt"

	c "github.com/y-hatano-github/coordin"
)

func main() {
	var m [26][26]string
	for i, rows := range m {
		for j := range rows {
			m[i][j] = "."
		}
	}

	lp, fp := c.Polygon(c.Points{
		{X: 12, Y: 2},
		{X: 16, Y: 6},
		{X: 22, Y: 6},
		{X: 18, Y: 10},
		{X: 22, Y: 15},
		{X: 12, Y: 12},
		{X: 3, Y: 15},
		{X: 7, Y: 10},
		{X: 3, Y: 6},
		{X: 8, Y: 6}})

	for _, p := range fp {
		m[p.X][p.Y] = "#"
	}
	for _, p := range lp {
		m[p.X][p.Y] = "*"
	}
	for i, rows := range m {
		for j := range rows {
			fmt.Print(m[j][i])
		}
		fmt.Println()
	}
}
