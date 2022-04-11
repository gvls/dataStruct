package picture

import "fmt"

// package picture

// SparseArray suit element <= 1/3 of total contain
func SparseArray() {
	sa := [][3]int{
		// line column value
		{0, 1, 2}, // it can be replace by struct
		{0, 1, 2},
		// ...
	}
	fmt.Println(sa)
}
