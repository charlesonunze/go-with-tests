package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("say hello, <Person Name> :)", func(t *testing.T) {
		const personName = "Charles"

		got := hello(personName)
		want := "Hello, " + personName

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello, world if an empty string is passed", func(t *testing.T) {
		const personName = ""

		got := hello(personName)
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})
}
