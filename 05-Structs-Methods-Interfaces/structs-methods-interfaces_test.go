package interfaces

import "testing"

func Test(t *testing.T) {
	rectangle := Rectangle{2.0, 3.0}

	t.Run("calculate perimeter of a rectangle", func(t *testing.T) {
		got := Perimeter(rectangle)
		want := 10.0

		if want != got {
			t.Errorf("want %.2f but got %.2f", want, got)
		}
	})

	t.Run("calculate area of a rectangle", func(t *testing.T) {
		got := Area(rectangle)
		want := 6.0

		if want != got {
			t.Errorf("want %.2f but got %.2f", want, got)
		}
	})
}
