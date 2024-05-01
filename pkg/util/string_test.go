package util

import (
	"testing"
)

func TestRandString(t *testing.T) {
	// Define test cases
	addTests := []struct {
		arg1     int
		expected string
	}{
		{10, "DE60733E50"},
		{10, "9E1DEE6F7A"},
	}

	// Run the tests
	for _, test := range addTests {
		output := RandString(test.arg1)
		if output != test.expected {
			t.Errorf("RandString(%d) = %v; expected %v", test.arg1, output, test.expected)
		}
	}
}

func BenchmarkRandString10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = RandString(10)
	}
}

func BenchmarkRandString100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = RandString(100)
	}
}

func BenchmarkRandString1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = RandString(1000)
	}
}
