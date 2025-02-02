package combin

import (
	"math/big"
	"slices"
	"testing"

	"github.com/akramarenkov/safe/intspec"
	"github.com/akramarenkov/seq"
	"github.com/stretchr/testify/require"
)

type intAlias []int

func (intAlias) Method() {}

func TestEveryZero(t *testing.T) {
	expected := [][]int{}

	testEvery(t, nil, expected)
	testEvery(t, []int{}, expected)
}

func TestEvery1(t *testing.T) {
	expected := [][]int{
		{1},
	}

	testEvery(t, seq.Linear(1, 1), expected)
}

func TestEveryAlias(t *testing.T) {
	source := intAlias{1}

	for combination := range Every(source) {
		combination.Method()
	}

	require.Equal(t, uint64(1), EveryQuantity(source).Uint64())
	require.Equal(t, uint64(1), EverySize(source))
}

func TestEvery21(t *testing.T) {
	expected := [][]int{
		{2},
		{1},
		{2, 1},
	}

	testEvery(t, seq.Linear(2, 1), expected)
}

func TestEvery321(t *testing.T) {
	expected := [][]int{
		{3},
		{2},
		{1},
		{2, 1},
		{3, 2},
		{3, 1},
		{3, 2, 1},
	}

	testEvery(t, seq.Linear(3, 1), expected)
}

func TestEvery4321(t *testing.T) {
	expected := [][]int{
		{4},
		{3},
		{2},
		{1},
		{2, 1},
		{3, 2},
		{3, 1},
		{4, 3},
		{4, 2},
		{4, 1},
		{3, 2, 1},
		{4, 3, 1},
		{4, 2, 1},
		{4, 3, 2},
		{4, 3, 2, 1},
	}

	testEvery(t, seq.Linear(4, 1), expected)
}

func TestEvery54321(t *testing.T) {
	expected := [][]int{
		{5},
		{4},
		{3},
		{2},
		{1},
		{2, 1},
		{3, 2},
		{3, 1},
		{4, 3},
		{4, 2},
		{4, 1},
		{5, 4},
		{5, 3},
		{5, 2},
		{5, 1},
		{3, 2, 1},
		{4, 3, 1},
		{4, 2, 1},
		{4, 3, 2},
		{4, 3, 2, 1},
		{5, 2, 1},
		{5, 3, 2},
		{5, 3, 1},
		{5, 4, 3},
		{5, 4, 2},
		{5, 4, 1},
		{5, 3, 2, 1},
		{5, 4, 3, 1},
		{5, 4, 2, 1},
		{5, 4, 3, 2},
		{5, 4, 3, 2, 1},
	}

	testEvery(t, seq.Linear(5, 1), expected)
}

func TestEvery654321(t *testing.T) {
	expected := [][]int{
		{6},
		{5},
		{4},
		{3},
		{2},
		{1},
		{2, 1},
		{3, 2},
		{3, 1},
		{4, 3},
		{4, 2},
		{4, 1},
		{5, 4},
		{5, 3},
		{5, 2},
		{5, 1},
		{6, 5},
		{6, 4},
		{6, 3},
		{6, 2},
		{6, 1},
		{3, 2, 1},
		{4, 3, 1},
		{4, 2, 1},
		{4, 3, 2},
		{4, 3, 2, 1},
		{5, 2, 1},
		{5, 3, 2},
		{5, 3, 1},
		{5, 4, 3},
		{5, 4, 2},
		{5, 4, 1},
		{6, 5, 4},
		{6, 5, 3},
		{6, 5, 2},
		{6, 5, 1},
		{6, 4, 3},
		{6, 4, 2},
		{6, 4, 1},
		{6, 3, 2},
		{6, 3, 1},
		{6, 2, 1},
		{5, 3, 2, 1},
		{5, 4, 3, 1},
		{5, 4, 2, 1},
		{5, 4, 3, 2},
		{6, 5, 4, 3},
		{6, 5, 4, 2},
		{6, 5, 4, 1},
		{6, 5, 3, 2},
		{6, 5, 3, 1},
		{6, 5, 2, 1},
		{6, 4, 3, 2},
		{6, 4, 3, 1},
		{6, 4, 2, 1},
		{6, 3, 2, 1},
		{5, 4, 3, 2, 1},
		{6, 5, 4, 3, 2},
		{6, 5, 4, 3, 1},
		{6, 5, 4, 2, 1},
		{6, 5, 3, 2, 1},
		{6, 4, 3, 2, 1},
		{6, 5, 4, 3, 2, 1},
	}

	testEvery(t, seq.Linear(6, 1), expected)
}

func testEvery(t *testing.T, source []int, expected [][]int) {
	combinations := make([][]int, 0, EverySize(source))

	for combination := range Every(source) {
		combinations = append(combinations, slices.Clone(combination))
	}

	require.ElementsMatch(t, combinations, expected)
}

func TestEveryBreak(*testing.T) {
	for range Every(seq.Linear(15, 1)) {
		break
	}
}

func TestEverySizeMatch(t *testing.T) {
	source := seq.Linear(15, 1)

	size := uint64(0)

	for range Every(source) {
		size++
	}

	require.Equal(t, EverySize(source), size)
}

func TestEverySequencing(t *testing.T) {
	for combination := range Every(seq.Linear(15, 1)) {
		require.IsDecreasing(t, combination, "combination: %v", combination)
	}

	for combination := range Every(seq.Linear(1, 15)) {
		require.IsIncreasing(t, combination, "combination: %v", combination)
	}
}

func TestEveryQuantity(t *testing.T) {
	require.Equal(t, uint64(0), EveryQuantity[int](nil).Uint64())
	require.Equal(t, uint64(0), EveryQuantity([]int{}).Uint64())
	require.Equal(t, uint64(1), EveryQuantity(seq.Linear(1, 1)).Uint64())
	require.Equal(t, uint64(3), EveryQuantity(seq.Linear(2, 1)).Uint64())
	require.Equal(t, uint64(7), EveryQuantity(seq.Linear(3, 1)).Uint64())
	require.Equal(t, uint64(15), EveryQuantity(seq.Linear(4, 1)).Uint64())
	require.Equal(t, uint64(31), EveryQuantity(seq.Linear(5, 1)).Uint64())
	require.Equal(t, uint64(63), EveryQuantity(seq.Linear(6, 1)).Uint64())
	require.Equal(t, uint64(127), EveryQuantity(seq.Linear(7, 1)).Uint64())
	require.Equal(t, uint64(255), EveryQuantity(seq.Linear(8, 1)).Uint64())
	require.Equal(t, uint64(511), EveryQuantity(seq.Linear(9, 1)).Uint64())
}

func TestEverySize(t *testing.T) {
	require.Equal(t, uint64(0), EverySize[int](nil))
	require.Equal(t, uint64(0), EverySize([]int{}))
	require.Equal(t, uint64(1), EverySize(seq.Linear(1, 1)))
	require.Equal(t, uint64(3), EverySize(seq.Linear(2, 1)))
	require.Equal(t, uint64(7), EverySize(seq.Linear(3, 1)))
	require.Equal(t, uint64(15), EverySize(seq.Linear(4, 1)))
	require.Equal(t, uint64(31), EverySize(seq.Linear(5, 1)))
	require.Equal(t, uint64(63), EverySize(seq.Linear(6, 1)))
	require.Equal(t, uint64(127), EverySize(seq.Linear(7, 1)))
	require.Equal(t, uint64(255), EverySize(seq.Linear(8, 1)))
	require.Equal(t, uint64(511), EverySize(seq.Linear(9, 1)))

	require.Equal(t, uint64(1<<63-1), EverySize(seq.Linear(63, 1)))
	require.Equal(t, uint64(intspec.MaxUint64), EverySize(seq.Linear(64, 1)))
}

func BenchmarkEveryQuantity(b *testing.B) {
	slice := seq.Linear(15, 1)

	var quantity *big.Int

	for range b.N {
		quantity = EveryQuantity(slice)
	}

	require.Equal(b, uint64(1<<15-1), quantity.Uint64())
}

func BenchmarkEverySize(b *testing.B) {
	slice := seq.Linear(15, 1)

	var size uint64

	for range b.N {
		size = EverySize(slice)
	}

	require.Equal(b, uint64(1<<15-1), size)
}

func BenchmarkReference(b *testing.B) {
	var number int

	for range b.N {
		for number = range 1<<15 - 1 {
			_ = number
		}
	}

	require.Equal(b, 1<<15-2, number)
}

func BenchmarkEvery(b *testing.B) {
	slice := seq.Linear(15, 1)

	var combination []int

	for range b.N {
		for combination = range Every(slice) {
			_ = combination
		}
	}

	require.NotNil(b, combination)
}
