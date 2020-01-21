package main

import "testing"

func TestHello(t *testing.T){
	got := Hello("mert")
	want := "Hello mert!"

	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}
