package coordin

import (
	"fmt"
	"math"
)

// Circle returns the coordinates of a circle.
// The First two parameters set the location, and the third and fourth parameters set radius of the x-axis and y-axis.
// Circle uses trigonometric functions to calculate coordinates.
// The First return value is perimeter points and the secound return value is fill points.
func Circle(cx, cy, rx, ry int) (Points, Points) {

	ps := Points{}
	s := 0.0
	if 360 <= (math.Max(float64(rx), float64(ry)) * 2 * 3.14) {
		s = 1
	} else {
		s = 360 / (math.Max(float64(rx), float64(ry)) * 2 * 3.14)
	}

	m := make(map[string]interface{})
	for i := 0.0; i <= 90; i += s {
		x := float64(rx) * math.Cos(float64(i)*3.14/180)
		y := float64(ry) * math.Sin(float64(i)*3.14/180)
		ps = append(ps, Point{X: cx - int(x), Y: cy - int(y)}) // 0 to 90 degrees
		ps = append(ps, Point{X: cx + int(x), Y: cy - int(y)}) // 90 to 180 degrees
		ps = append(ps, Point{X: cx + int(x), Y: cy + int(y)}) // 180 to 270 degrees
		ps = append(ps, Point{X: cx - int(x), Y: cy + int(y)}) // 270 to 360 degrees

		m[fmt.Sprintf("%d,%d", cx-int(x), cy-int(y))] = ""
		m[fmt.Sprintf("%d,%d", cx+int(x), cy-int(y))] = ""
		m[fmt.Sprintf("%d,%d", cx+int(x), cy+int(y))] = ""
		m[fmt.Sprintf("%d,%d", cx-int(x), cy+int(y))] = ""
	}

	var fp *Points = &Points{}
	fillc(cx, cy, m, fp)
	return ps, *fp
}

// fillc is a function that calculates data for fill.
func fillc(x, y int, m map[string]interface{}, fp *Points) {
	if _, ok := m[fmt.Sprintf("%d,%d", x, y)]; !ok {
		m[fmt.Sprintf("%d,%d", x, y)] = ""
		*fp = append(*fp, Point{X: x, Y: y})
		fillc(x, y-1, m, fp)
		fillc(x, y+1, m, fp)
		fillc(x-1, y, m, fp)
		fillc(x+1, y, m, fp)
	}
}

// Circle2 returns the coordinates of a circle.
// The First two parameters set the location, and the third and fourth parameters set radius of the x-axis and y-axis.
// Circle2 calculates the distance from the center point to the circumference without using trigonometric.
// The First return value is perimeter points and the secound return value is fill points.
func Circle2(cx, cy, rx, ry int) (Points, Points) {
	ps := Points{}
	lp := Points{}
	fp := Points{}
	m := make(map[string]interface{})

	for x := cx - rx; x <= cx+rx; x++ {
		for y := cy - ry; y <= cy+ry; y++ {
			dx := math.Abs(float64(x - cx))
			dy := math.Abs(float64(y - cy))

			if (dx*dx/float64(rx*rx) + dy*dy/float64(ry*ry)) <= 1 {
				m[fmt.Sprintf("%d,%d", x, y)] = ""
				ps = append(ps, Point{X: x, Y: y})
			}
		}
	}

	for _, p := range ps {
		_, left := m[fmt.Sprintf("%d,%d", p.X-1, p.Y)]
		_, right := m[fmt.Sprintf("%d,%d", p.X+1, p.Y)]
		_, above := m[fmt.Sprintf("%d,%d", p.X, p.Y-1)]
		_, below := m[fmt.Sprintf("%d,%d", p.X, p.Y+1)]

		if left && right && above && below {
			fp = append(fp, p)
		} else {
			lp = append(lp, p)
		}
	}
	return lp, fp
}

// Circled returns the coordinates of the dashed circle.
// The First two parameters set the location, and the third and fourth parameters set radius of the x-axis and y-axis and the fifth parameter set interval of the dash spacing.
// The calculation method is same as `Circle`.
func Circled(cx, cy, rx, ry, s int) Points {
	ps := Points{}
	ds := 0.0
	if s <= 1 {
		if 360 <= (math.Max(float64(rx), float64(ry)) * 2 * 3.14) {
			ds = 1
		} else {
			ds = 360 / (math.Max(float64(rx), float64(ry)) * 2 * 3.14)
		}
	}

	for i := 0.0; i <= 90; i += ds {
		x := float64(rx) * math.Cos(float64(i)*3.14/180)
		y := float64(ry) * math.Sin(float64(i)*3.14/180)
		ps = append(ps, Point{X: cx - int(x), Y: cy - int(y)}) // 0 to 90 degrees
		ps = append(ps, Point{X: cx + int(x), Y: cy - int(y)}) // 90 to 180 degrees
		ps = append(ps, Point{X: cx + int(x), Y: cy + int(y)}) // 180 to 270 degrees
		ps = append(ps, Point{X: cx - int(x), Y: cy + int(y)}) // 270 to 360 degrees
	}

	return ps
}

// Arc returns the coordinates of a arc.
// The First two parameters set the location, and the third and fourth parameters set radius of the x-axis and y-axis and the fifth and sixth parameters set the start and end degree.
func Arc(cx, cy, rx, ry, sd, ed int) Points {

	ps := Points{}
	s := 0.0
	if 360 <= (math.Max(float64(rx), float64(ry)) * 2 * 3.14) {
		s = 1
	} else {
		s = 360 / (math.Max(float64(rx), float64(ry)) * 2 * 3.14)
	}

	a := math.Min(float64(sd), float64(ed))
	b := math.Max(float64(sd), float64(ed))
	for i := a; i <= b; i += s {
		x := float64(rx) * math.Cos(float64(i)*3.14/180)
		y := float64(ry) * math.Sin(float64(i)*3.14/180)
		ps = append(ps, Point{X: cx + int(x), Y: cy + int(y)})
	}

	return ps
}
