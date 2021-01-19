package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	preSum := A[0]
	max := A[0]
	for i := 1; i < len(A); i++ {
		curSum := preSum + A[i]
		if A[i] < curSum {
			preSum = curSum
		} else {
			preSum = A[i]
		}

		if preSum > max {
			max = preSum
		}
	}
	return max
}
