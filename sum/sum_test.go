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

func TestSumAll(t *testing.T) {

	numbers := []int{1, 2, 3}
	numbers1 := []int{2, 3, 4, 5}

	got := SumALL(numbers, numbers1)
	want := []int{6, 14}

	if !slices.Equal(got, want) {
		t.Errorf("Got %d, want %d, from array %v and array %v", got, want, numbers, numbers1)
	}

}

func TestSumAllTail(t *testing.T) {

	assertCorrectMessage := func (t testing.TB, got, want []int) {
	t.Helper()
	if !slices.Equal(got, want) {
		t.Errorf("got %d, want %d", got, want)
	}
}

	t.Run("Make the sum of some slices", func(t *testing.T) {
		got := SumAllTail([]int{0, 2}, []int{2, 9, 3})
		want := []int{2, 12}

		assertCorrectMessage(t, got, want)

	})

	t.Run("Safely sum empty slices", func(t *testing.T) {
		got := SumAllTail([]int{}, []int{2, 9, 3})
		want := []int{0, 12}
		assertCorrectMessage(t, got, want)
	})

}


