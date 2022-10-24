package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 2)
	want := "aa"

	if got != want {
		t.Errorf("got: %q want: %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("b", 3)
	}

}

func ExampleRepeat() {
	repeated := Repeat("c", 4)
	fmt.Println(repeated)
	// Output: cccc
}
