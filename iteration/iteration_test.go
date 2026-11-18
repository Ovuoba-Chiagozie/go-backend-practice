package iteration

import "testing"

func TestIteration(t *testing.T) {

	got := Repeat("a")
	expected := "aaaaa"

	if got != expected {
	 t.Errorf("got %q expected %q", got, expected)
	}

}

func BenchmarkRepeat(b *testing.B) {
 
	for b.Loop() {
		Repeat("a")
	}
}
