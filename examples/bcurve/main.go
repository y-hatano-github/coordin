package main

import (
	"fmt"

	c "github.com/yoshihicode/coordin"
)

func main() {
	fmt.Print("\033[2J")
	ps := c.Points{
		{X: 0, Y: 15},
		{X: 8, Y: 5},
		{X: 18, Y: 5},
		{X: 24, Y: 15},
		{X: 40, Y: 15},
		{X: 45, Y: 10},
	}

	bps := c.BCurve(ps, 0.7)

	rendering(bps, "+")
	fmt.Printf("\033[%v;%vH", 17, 0)

}

func rendering(ps c.Points, s string) {
	for _, p := range ps {
		fmt.Printf("\033[%v;%vH"+s, p.Y, p.X)
	}
}
