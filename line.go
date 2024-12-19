package coordin

import (
	"fmt"
	"math"
)

// Line returns the coordinates of a line between two points.
func Line(p1 Point, p2 Point) Points {

	var dx, dy, sx, sy int
	x1 := p1.X
	y1 := p1.Y
	x2 := p2.X
	y2 := p2.Y

	if x2 > x1 {
		sx = 1
	} else {
		sx = -1
	}
	if x2 > x1 {
		dx = x2 - x1
	} else {
		dx = x1 - x2
	}
	if y2 > y1 {
		sy = 1
	} else {
		sy = -1
	}
	if y2 > y1 {
		dy = y2 - y1
	} else {
		dy = y1 - y2
	}

	ps := []Point{}

	x := x1
	y := y1

	if dx >= dy {
		e := -dx
		for i := 0; i <= dx; i++ {
			ps = append(ps, Point{X: x, Y: y})
			x += sx
			e += 2 * dy
			if e >= 0 {
				y += sy
				e -= 2 * dx
			}
		}

	} else {
		e := -dy
		for i := 0; i <= dy; i++ {
			ps = append(ps, Point{X: x, Y: y})
			y += sy
			e += 2 * dx
			if e >= 0 {
				x += sx
				e -= 2 * dy
			}
		}
	}

	return ps
}

// Rect returns the coordinates of a rectangle.
// First return value is perimeter points, secound return value is fill points.
func Rect(p1 Point, p2 Point) (Points, Points) {
	p3 := Point{p1.X, p2.Y}
	p4 := Point{p2.X, p1.Y}

	bx := Points{}
	bx = append(bx, Line(p1, p4)...)
	bx = append(bx, Line(p3, p2)...)
	bx = append(bx, Line(p1, p3)...)
	bx = append(bx, Line(p4, p2)...)

	minX := math.Min(float64(p1.X), float64(p2.X))
	minY := math.Min(float64(p1.Y), float64(p2.Y))

	fp := Points{}
	for x := minX + 1; x < math.Abs(float64(p1.X-p2.X)); x++ {
		for y := minY + 1; x < math.Abs(float64(p1.Y-p2.Y)); y++ {
			fp = append(fp, Point{X: int(x), Y: int(y)})
		}
	}
	return bx, fp
}

// Polyline returns coordinates connecting the given points with a line.
// There is no line connecting the last point and the first point.
func Polyline(ps Points) Points {
	pl := Points{}

	for i := 0; i < len(ps)-1; i++ {
		pl = append(pl, Line(Point{ps[i].X, ps[i].Y}, Point{ps[i+1].X, ps[i+1].Y})...)
	}

	return pl
}

// Polygon returns the coordinates connecting the given points and the line connecting the last point and the first point.
// The First return value is perimeter points, and the second is fill points.
func Polygon(ps Points) (Points, Points) {
	pl := Points{}

	for i := 0; i < len(ps)-1; i++ {
		pl = append(pl, Line(Point{ps[i].X, ps[i].Y}, Point{ps[i+1].X, ps[i+1].Y})...)
	}
	pl = append(pl, Line(Point{ps[len(ps)-1].X, ps[len(ps)-1].Y}, Point{ps[0].X, ps[0].Y})...)

	m := make(map[string]interface{})
	for _, p := range pl {
		m[fmt.Sprintf("%d,%d", p.X, p.Y)] = ""
	}
	var x, y int
	for _, p := range ps {
		x = x + p.X
		y = y + p.Y
	}

	var fp *Points = &Points{}
	fillp(int(x/(len(ps)-1)), int(y/(len(ps)-1)), m, fp)

	return pl, *fp
}

// fillp is a function that calculates data for fill.
func fillp(x, y int, m map[string]interface{}, fp *Points) {
	if _, ok := m[fmt.Sprintf("%d,%d", x, y)]; !ok {
		m[fmt.Sprintf("%d,%d", x, y)] = ""
		*fp = append(*fp, Point{X: x, Y: y})
		fillp(x, y-1, m, fp)
		fillp(x, y+1, m, fp)
		fillp(x-1, y, m, fp)
		fillp(x+1, y, m, fp)
	}
}
