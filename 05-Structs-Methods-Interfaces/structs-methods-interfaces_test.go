package interfaces

import "testing"

func Test(t *testing.T) {
	checkArea := func(t *testing.T, s Shape, want float64) {
		t.Helper()
		got := s.Area()

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("calculate area of a rectangle", func(t *testing.T) {
		rectangle := Rectangle{2.0, 3.0}
		want := 6.0
		checkArea(t, rectangle, want)
	})

	t.Run("calculate area of a circle", func(t *testing.T) {
		circle := Circle{2.0}
		want := 12.566370614359172
		checkArea(t, circle, want)
	})
}
