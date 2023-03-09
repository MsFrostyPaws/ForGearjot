package primes

import (
	"math"
)

func IsPrime(num int) bool {
	if num <= 3 {
		return num > 1
	}
	if num%2 == 0 || num%3 == 0 {
		return false
	}
	limit := int(math.Pow(float64(num), 0.5))

	for i := 5; i <= limit+1; i += 6 {
		if num%i == 0 || num%(i+2) == 0 {
			return false
		}
	}
	return true
}
