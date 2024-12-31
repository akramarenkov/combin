package combin_test

import (
	"fmt"

	"github.com/akramarenkov/combin"
)

func ExampleEvery() {
	for combination := range combin.Every([]int{4, 3, 2, 1}) {
		fmt.Println(combination)
	}
	// Output:
	// [4]
	// [4 3]
	// [4 3 2]
	// [4 3 2 1]
	// [4 3 1]
	// [4 2]
	// [4 2 1]
	// [4 1]
	// [3]
	// [3 2]
	// [3 2 1]
	// [3 1]
	// [2]
	// [2 1]
	// [1]
}
