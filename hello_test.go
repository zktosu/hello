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
		got := Hello("mert","english")
		want := "Hello mert!"
		assertCorrectMessage(t,got,want)
	})

	t.Run("No named greeting:", func(t *testing.T){
		got := Hello("","english")
		want := "Hello World!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("greetings in Turkish", func(t *testing.T){
		got := Hello("çökelek", "turkish")
		want := "Merhaba çökelek!"
		assertCorrectMessage(t, got, want)
	})
	
	t.Run("greetings in french", func(t *testing.T){
		got := Hello("sebastian","french")
		want := "Bonjour sebastian!"
		assertCorrectMessage(t, got, want)
	})

}
