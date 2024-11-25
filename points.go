package coordin

// Points is array of Point.
type Points []Point

// Point is a structure of XY coordinates.
// {X:0, Y:0} indicates the top left edge.
type Point struct {
	// X coordinate
	X int
	// Y coordinate
	Y int
}
