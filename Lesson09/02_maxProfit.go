package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	if len(A) < 2 {
		return 0
	}
	minPrice := A[0]
	maxProfit := 0
	for i := 1; i < len(A); i++ {
		curProfit := A[i] - minPrice
		if curProfit > 0 {
			if curProfit > maxProfit {
				maxProfit = curProfit
			}
		} else {
			minPrice = A[i]
		}
	}
	return maxProfit
}
