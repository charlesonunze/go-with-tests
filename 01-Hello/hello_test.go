package main

import "testing"

func TestHello(t *testing.T) {
	got := hello("charles")
	want := "hello, charles"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
