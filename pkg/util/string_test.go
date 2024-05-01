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

func TestRandHex(t *testing.T) {
	// Define test cases
	addTests := []struct {
		arg1     int
		expected string
	}{
		{10, "B9F41DB402"},
		{10, "3F9EDA8FE1"},
	}

	// Run the tests
	for _, test := range addTests {
		output := RandHex(test.arg1)
		if output != test.expected {
			t.Errorf("RandHex(%d) = %v; expected %v", test.arg1, output, test.expected)
		}
	}
}
func BenchmarkRandHex10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = RandHex(10)
	}
}

func BenchmarkRandHex100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = RandHex(100)
	}
}

func BenchmarkRandHex1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = RandHex(1000)
	}
}
