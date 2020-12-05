package interfaces

// Rectangle struct
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter function
func Perimeter(r Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}

// Area function
func Area(r Rectangle) float64 {
	return r.Width * r.Height
}
