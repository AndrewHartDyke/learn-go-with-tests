package main

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	/// ...setup...
	for b.Loop() {
		Repeat("a", 5)
	}
	/// ...cleanup...
}

func ExampleRepeat() {
	repeated := Repeat("9", 20)
	fmt.Println(repeated)
	// Output: 99999999999999999999
}
