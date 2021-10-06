package main

import "testing"

func TestHello(t *testing.T) {
	asserCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Robin", "")
		want := "Hello, Robin"
		asserCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, world' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		asserCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Robin", "Spanish")
		want := "Hola, Robin"
		asserCorrectMessage(t, got, want)
	})
}
