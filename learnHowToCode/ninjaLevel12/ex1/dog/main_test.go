package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	x := Years(7)
	if x != 49 {
		t.Error("Expected: ", 49, "Got: ", x)
	}
}

func ExampleYears() {
	fmt.Println(Years(7))
	// Output:
	// 49
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(7)
	}
}

func TestYearsTwo(t *testing.T) {
	x := Years(7)
	if x != 49 {
		t.Error("Expected: ", 49, "Got: ", x)
	}
}

func ExampleYearsTwo() {
	fmt.Println(Years(7))
	// Output:
	// 49
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(7)
	}
}
