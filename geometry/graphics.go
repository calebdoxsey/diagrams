package geometry

import "math"

// A Point is an {x,y} coordinate.
type Point struct {
	X, Y float64
}

// Translate translates a point by the given amount.
func (p Point) Translate(dx, dy float64) Point {
	return Point{p.X + dx, p.Y + dy}
}

// At returns a point for the given {x, y} coordinates.
func At(x, y float64) Point {
	return Point{X: x, Y: y}
}

// A Positionable object is one that has an {x,y} position.
type Positionable interface {
	GetPosition() Point
	SetPosition(Point)
}

// A Line is a line between two points.
type Line [2]Point

// Length returns the length of the line.
func (l Line) Length() float64 {
	dx, dy := l[0].X-l[1].X, l[0].Y-l[1].Y
	return math.Sqrt(dx*dx + dy*dy)
}

// HasVisibility is an interface for objects that have a visibility property.
type HasVisibility interface {
	GetVisibility() float64
	SetVisibility(float64)
}
