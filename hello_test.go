package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("saying hello to people", func(t *testing.T) {

		got := Hello("Chris", "")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello when its an empty string", func(t *testing.T) {

		got := Hello("","")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)

	})

	t.Run("In Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		assertCorrectMessage(t,got,want)
	})

	t.Run("In French", func(t *testing.T) {
		got := Hello("Christopher", "French")
		want := "Bonjour, Christopher"

		assertCorrectMessage(t,got,want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if want != got {
		t.Errorf("got %q want %q", got, want)
	}
}
