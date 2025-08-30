package iterators

import (
	"iter"
)

func Primes(seq iter.Seq[int]) iter.Seq[int] {
	return func(yield func(int) bool) {
		for n := range seq {
			if isPrime(n) {
				if !yield(n) {
					return
				}
			}
		}
	}
}

func NonZeroIntegers(take int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := range take {
			if !yield(i + 1) {
				return
			}
		}
	}
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	limit := n / 2
	for i := 2; i <= limit; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
