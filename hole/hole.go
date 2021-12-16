package hole

import (
	"math/rand"

	"golang.org/x/exp/constraints"
)

// Doesn't handle any special cases, will be in the stdlib/x one day.
func max[T constraints.Ordered](x, y T) T {
	if x > y {
		return x
	}
	return y
}

// Doesn't handle any special cases, will be in the stdlib/x one day.
func min[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}

// Returning the slice is a convenience, the shuffle is still in-place.
func shuffle[E any](x []E) []E {
	rand.Shuffle(len(x), func(i, j int) { x[i], x[j] = x[j], x[i] })
	return x
}
