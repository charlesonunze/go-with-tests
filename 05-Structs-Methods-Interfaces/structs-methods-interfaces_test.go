package interfaces

import "testing"

func Test(t *testing.T) {
	rectangle := Rectangle{2.0, 3.0}
	circle := Circle{2.0}

	t.Run("calculate area of a rectangle", func(t *testing.T) {
		got := rectangle.Area()
		want := 6.0

		if want != got {
			t.Errorf("want %.2f but got %.2f", want, got)
		}
	})

	t.Run("calculate area of a circle", func(t *testing.T) {
		got := circle.Area()
		want := 12.566370614359172

		if want != got {
			t.Errorf("want %f but got %f", want, got)
		}
	})
}
