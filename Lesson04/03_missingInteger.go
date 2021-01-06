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
	result := 1
	if A[len(A)-1] < 1 {
		return result
	}

	index := 0
	for A[index] < 1 {
		index++
	}
	if A[index] > 1 {
		return result
	}

	result = A[index]
	for index < len(A) {
		if A[index] == result {
			result++
		} else if A[index] != A[index-1] {
			return result
		}
		index++
	}
	return result
}
