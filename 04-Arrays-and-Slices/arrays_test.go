package slices

import "testing"

func Test(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		got := Sum(input)
		want := 15

		if got != want {
			t.Errorf("want %d but got %d given %v", want, got, input)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		input := []int{1, 2, 3}
		got := Sum(input)
		want := 6

		if got != want {
			t.Errorf("want %d but got %d given %v", want, got, input)
		}
	})
}
