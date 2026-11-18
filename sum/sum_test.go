package main

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {


	t.Run("Collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("Got %d, want %d, from array %v", got, want, numbers)
		}
	})

}

func TestSumAll(t * testing.T) {

	numbers := []int{1,2,3}
	numbers1 := []int{2,3,4,5}

	got := SumALL(numbers,numbers1)
    want := []int{6,14}

	if !slices.Equal(got,want) {
		t.Errorf("Got %d, want %d, from array %v and array %v", got,want,numbers,numbers1)
	}

}
