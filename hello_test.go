package main

import "testing"

func TestHello(t *testing.T) {

	got := Hello("Chris")
	want := "Hello, Chris"

	if want != got {
		t.Errorf("got %q want %q", got, want)
	}

}