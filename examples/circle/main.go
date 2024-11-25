package main

import (
	"fmt"

	c "github.com/y-hatano-github/coordin"
)

func main() {
	fmt.Print("\033[2J")

	circle(c.Circle, 22, 12, 20, 10, "#", "+")
	circle(c.Circle2, 64, 12, 20, 10, "+", "O")

	ap := c.Arc(105, 12, 20, 10, 20, 160)
	rendering(ap, "$")

	fmt.Printf("\033[%v;%vH", 23, 1)
}

func circle(f func(int, int, int, int) (c.Points, c.Points), x, y, rx, ry int, sl, sf string) {

	cl, cf := f(x, y, rx, ry)

	rendering(cl, sl)
	rendering(cf, sf)

}

func rendering(ps c.Points, s string) {
	for _, p := range ps {
		fmt.Printf("\033[%v;%vH"+s, p.Y, p.X)
	}
}
