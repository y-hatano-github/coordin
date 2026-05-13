package coordin

import (
	"math"
	"sort"
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

	minX := int(math.Min(float64(p1.X), float64(p2.X)))
	maxX := int(math.Max(float64(p1.X), float64(p2.X)))
	minY := int(math.Min(float64(p1.Y), float64(p2.Y)))
	maxY := int(math.Max(float64(p1.Y), float64(p2.Y)))

	fp := Points{}
	for x := minX + 1; x < maxX; x++ {
		for y := minY + 1; y < maxY; y++ {
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

	fp := fillp(ps)
	return pl, fp
}

// fillp is a function that calculates data for fill.
func fillp(ps Points) Points {
	rp := Points{}

	ymin, ymax := ps[0].Y, ps[0].Y
	for _, v := range ps {
		if v.Y < ymin {
			ymin = v.Y
		}
		if v.Y > ymax {
			ymax = v.Y
		}
	}

	for sy := ymin; sy <= ymax; sy++ {
		var xs []int

		for i := 0; i < len(ps); i++ {
			j := (i + 1) % len(ps)

			y1, y2 := ps[i].Y, ps[j].Y
			x1, x2 := ps[i].X, ps[j].X

			if (y1 <= sy && y2 > sy) ||
				(y2 <= sy && y1 > sy) {

				fx := float64(x1) +
					float64(sy-y1)*float64(x2-x1)/float64(y2-y1)

				xs = append(xs, int(math.Round(fx)))
			}
		}

		sort.Ints(xs)

		for i := 0; i+1 < len(xs); i += 2 {
			for sx := xs[i] + 1; sx < xs[i+1]; sx++ {
				rp = append(rp, Point{X: sx, Y: sy})
			}
		}
	}

	return rp
}
