package slices

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		got := Sum(input)
		want := 15

		if got != want {
			t.Errorf("want %d but got %d given %v", want, got, input)
		}
	})

	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %d but got %d given", want, got)
		}
	}

	t.Run("sum all", func(t *testing.T) {
		input1 := []int{1, 2}
		input2 := []int{3, 4}
		got := SumAll(input1, input2)
		want := []int{3, 7}

		checkSums(t, got, want)
	})

	t.Run("sum all tails", func(t *testing.T) {
		input1 := []int{1, 2, 3}
		input2 := []int{3, 4, 5}
		got := SumAllTails(input1, input2)
		want := []int{5, 9}

		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		input1 := []int{}
		input2 := []int{3, 4, 5}
		got := SumAllTails(input1, input2)
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}
