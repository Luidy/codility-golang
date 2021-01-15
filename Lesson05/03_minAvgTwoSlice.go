package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	if len(A) == 2 {
		return 0
	}

	min := float64(A[0]+A[1]) / 2
	minIndex := 0
	for i := 2; i < len(A); i++ {
		avg := float64(A[i-2]+A[i-1]+A[i]) / 3
		if avg < min {
			minIndex = i - 2
			min = avg
		}
		avg = float64(A[i-1]+A[i]) / 2
		if avg < min {
			minIndex = i - 1
			min = avg
		}
	}
	return minIndex
}
