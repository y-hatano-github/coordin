package coordin

// BCurve returns the coordinates of Bézier curve.
// The first and last indices of the first parameter(Points) are endpoints, and the others are control points.
// The second parameter(float64) is curvature.
func BCurve(ps Points, t float64) Points {

	var cps *Points = &Points{}

	getBCurnvePoints(ps, t, cps)

	return Polyline(*cps)
}

// getBCurnvePoints
func getBCurnvePoints(ps Points, t float64, cps *Points) {
	*cps = append(*cps, ps[0])
	if len(ps) > 1 {
		newPs := Points{}
		for i := 0; i < len(ps)-1; i++ {
			x := (1-t)*float64(ps[i].X) + t*float64(ps[i+1].X)
			y := (1-t)*float64(ps[i].Y) + t*float64(ps[i+1].Y)
			newPs = append(newPs, Point{X: int(x), Y: int(y)})
		}
		getBCurnvePoints(newPs, t, cps)
		*cps = append(*cps, ps[len(ps)-1])
	}
}
