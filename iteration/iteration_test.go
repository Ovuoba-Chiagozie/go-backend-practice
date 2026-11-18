package iteration

import (
	"fmt"
	"testing"
)

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

func ExampleRepeat() {
	repeated := Repeat("a")
    fmt.Println(repeated)
	// Output: "aaaaa"
}