package integer

import(
	"testing"
	"fmt"
)
func TestAdder(t *testing.T){
	sum := Add(2,3)
	want := 5

	if sum != want {
		t.Errorf("expected value %d, but got %d", sum, want)
	}
}
func ExampleAdd(){
	sum := Add(1,6)
	fmt.Println(sum)
	// Output: 7
}
