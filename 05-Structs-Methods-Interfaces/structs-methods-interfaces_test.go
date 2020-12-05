package interfaces

import "testing"

func Test(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Square{2.0, 2.0}, 4.0},
		{Triangle{4.0, 5.0}, 10.0},
		{Rectangle{2.0, 3.0}, 6.0},
		{Circle{2.0}, 12.566370614359172},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}
