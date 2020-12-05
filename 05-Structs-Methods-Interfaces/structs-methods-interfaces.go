package interfaces

import (
	"math"
)

// Shape interface
type Shape interface {
	Area() float64
}

// Rectangle struct
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter function
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Area function
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle struct
type Circle struct {
	Radius float64
}

// Area function
func (c Circle) Area() float64 {
	return math.Pi * (c.Radius * c.Radius)
}

// Square struct
type Square struct {
	Width  float64
	Height float64
}

// Area function
func (s Square) Area() float64 {
	return s.Width * s.Height
}

// Triangle struct
type Triangle struct {
	Base   float64
	Height float64
}

// Area function
func (t Triangle) Area() float64 {
	return (0.5 * t.Base) * t.Height
}
