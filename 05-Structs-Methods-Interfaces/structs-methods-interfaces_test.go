package interfaces

import "testing"

func Test(t *testing.T) {
	got := Perimeter(2.0, 3.0)
	want := 10.0

	if want != got {
		t.Errorf("want %.2f but got %.2f", want, got)
	}
}
