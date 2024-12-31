// Library that allows you to drawn combinations.
package combin

import (
	"iter"
	"math/big"

	"github.com/akramarenkov/safe"
	"github.com/akramarenkov/safe/intspec"
)

// A range iterator over every possible combinations (of length from 1 to the length
// of the specified slice) of slice elements without repeating elements within a single
// combination.
//
// The returned slice of combination is valid only for current iteration of the loop.
func Every[Type any](source []Type) iter.Seq[[]Type] {
	iterator := func(yield func([]Type) bool) {
		if len(source) == 0 {
			return
		}

		combination := make([]Type, len(source))
		shifts := make([]int, len(source))
		ids := make([]int, len(source))

		for level := 0; level != -1; {
			remainder := source[level+shifts[level]+ids[level]:]

			combination[level] = remainder[0]

			if !yield(combination[:level+1]) {
				return
			}

			if len(remainder) == 1 {
				ids[level] = 0
				level--

				continue
			}

			shifts[level+1] = shifts[level] + ids[level]
			ids[level]++
			level++
		}
	}

	return iterator
}

// Returns the quantity of combinations that can be obtained from the [Every] function.
//
// If the number of combinations for n elements is m, then for n+1 elements the
// number of combinations is m + m + 1. It is easy to see that such an increment
// corresponds to the function 2^n - 1.
func EveryQuantity[Type any](source []Type) *big.Int {
	quantity := new(big.Int).Sub(
		new(big.Int).Lsh(
			big.NewInt(1),
			uint(len(source)),
		),
		big.NewInt(1),
	)

	return quantity
}

// Returns the size of combinations that can be obtained from the [Every] function.
//
// The returned value is intended to be used in the make call as the size parameter so
// is truncated to the maximum value for uint64 if the calculated value exceeds it.
func EverySize[Type any](source []Type) uint64 {
	pow, err := safe.Shift[uint64](1, len(source))
	if err != nil {
		return intspec.MaxUint64
	}

	return pow - 1
}
