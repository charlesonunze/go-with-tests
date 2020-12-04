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

	t.Run("sum all", func(t *testing.T) {
		input1 := []int{1, 2}
		input2 := []int{3, 4}
		got := SumAll(input1, input2)
		want := []int{3, 7}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %d but got %d given %v and %v", want, got, input1, input2)
		}
	})

	t.Run("sum all tails", func(t *testing.T) {
		input1 := []int{1, 2, 3}
		input2 := []int{3, 4, 5}
		got := SumAllTails(input1, input2)
		want := []int{5, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %d but got %d given %v and %v", want, got, input1, input2)
		}
	})
}
