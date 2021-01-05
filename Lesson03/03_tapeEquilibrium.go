package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

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
	right := 0
	for i := 1; i < len(A); i++ {
		right = right + A[i]
	}
	left := A[0]
	diff := abs(left - right)
	for i := 1; i < len(A)-1; i++ {
		left = left + A[i]
		right = right - A[i]
		curDiff := abs(left - right)
		if curDiff < diff {
			diff = curDiff
		}
	}
	return diff
}
