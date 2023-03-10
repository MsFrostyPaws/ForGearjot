package primes

import "testing"

func TestIsPrime(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{"negative numbers", -99, false},
		{"0", 0, false},
		{"2", 2, true},
		{"even number", 4, false},
		{"prime number", 11, true},
		{"6", 6, false},
		{"1", 1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if IsPrime(tt.input) != tt.expected {
				t.Errorf("expected (%d) to be %v, but it should have been %v", tt.input, tt.expected, !tt.expected)
			}
		})
	}
}

func Benchmark(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsPrime(n)
	}
}
