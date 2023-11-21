package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	x := Years(3)
	if x != 21 {
		t.Error("got", x, "want", 21)
	}
}

func ExampleYears() {
	fmt.Println(Years(5))
	//Output:
	//35
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(20)
	}
}
