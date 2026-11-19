package structs

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(10, 20)
	want := 60.00

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	got := Area(10,20)
	want := 200.0

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}