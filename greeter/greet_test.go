package greeter

import "testing"

func TestGreet(t *testing.T){
	assertTest := func(t *testing.T, got, want string){
		t.Helper()
		if got != want {
			t.Errorf("Got %q, Want: %q",got, want)
		}
	}
	t.Run("This one test without parameter", func(t *testing.T){
		got := Greet()
		want := "Hello"
		assertTest(t, got, want)
	})
	t.Run("parametreli hali", func(t *testing.T){
		got := GreetName("Durmuş")
		want:="Hello, Durmuş!"
		assertTest(t, got, want)
	})
}

