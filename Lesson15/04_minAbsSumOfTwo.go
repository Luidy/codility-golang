package solution

// you can also use imports, for example:
// import "fmt"
// import "os"
import "sort"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Solution(A []int) int {
	// write your code in Go 1.4
	sort.Ints(A)
	left, right := 0, len(A)-1
	min := abs(A[left] + A[right])

	for left <= right {
		cur := A[left] + A[right]
		if cur == 0 {
			return cur
		}

		if min > abs(cur) {
			min = abs(cur)
		}

		if cur < 0 {
			left++
		} else {
			right--
		}
	}
	return min
}
