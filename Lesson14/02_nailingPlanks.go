package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int, B []int, C []int) int {
	// write your code in Go 1.4
	min := 1
	max := len(C)
	nail := make([]int, len(C)*2+1)
	total := -1
	for min <= max {
		miss := false
		mid := (min + max) / 2
		for i := 0; i < len(nail); i++ {
			nail[i] = 0
		}
		for i := 0; i < mid; i++ {
			nail[C[i]]++
		}
		for i := 1; i < len(nail); i++ {
			nail[i] = nail[i] + nail[i-1]
		}
		for i := 0; i < len(A); i++ {
			if nail[A[i]-1] == nail[B[i]] {
				miss = true
			}
		}
		if miss {
			min = mid + 1
		} else {
			max = mid - 1
			total = mid
		}
	}
	return total
}
