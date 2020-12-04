package slices

import "testing"

func Test(t *testing.T) {
	input := [5]int{1, 2, 3, 4, 5}
	got := Sum(input)
	want := 15

	if got != want {
		t.Errorf("want %d but got %d given %v", want, got, input)
	}
}
