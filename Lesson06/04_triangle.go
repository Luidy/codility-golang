package solution

// you can also use imports, for example:
// import "fmt"
// import "os"
import "sort"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	sort.Ints(A)
	for i := len(A) - 1; i > 1; i-- {
		if A[i] < A[i-1]+A[i-2] {
			return 1
		}
	}
	return 0
}
