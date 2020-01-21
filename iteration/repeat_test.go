package iteration

import (
	"testing"
	"fmt"
)

func TestRepeat(t *testing.T){
	got := Repeat("a",5)
	want := "aaaaa"

	if got != want {
		t.Errorf("got %q but expected %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B){
	for i := 0; i < b.N; i++ {
		Repeat("a",5)
	}
}
func ExampleRepeat(){
	repeated := Repeat("b",3)
	fmt.Println(repeated)
	//Output: bbb
}
