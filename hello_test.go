package main

import "testing"

func TestHello(t *testing.T){

	assertCorrectMessage := func(t *testing.T, got, want string){
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q",got, want)
		}
	}

	t.Run("Say hello to people", func(t *testing.T){
		got := Hello("mert")
		want := "Hello mert!"
		assertCorrectMessage(t,got,want)
	})

	t.Run("No named greeting:", func(t *testing.T){
		got := Hello("")
		want := "Hello World!"

		assertCorrectMessage(t, got, want)
	})
}
