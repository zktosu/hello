package main

import "testing"

func TestHello(t *testing.T){

	t.Run("Say hello to people", func(t *testing.T){
		got := Hello("mert")
		want := "Hello mert!"
		if got != want {
			t.Errorf("Got %q want %q", got, want)
		}
	})

	t.Run("No named greeting:", func(t *testing.T){
		got := Hello("")
		want := "Hello World!"

		if got != want {
			t.Errorf("got %q want %q",got,want)
		}
	})
}
