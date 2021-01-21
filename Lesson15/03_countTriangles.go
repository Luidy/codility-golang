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
	result := 0
	for p := 0; p < len(A)-2; p++ {
		r := p + 2
		for q := p + 1; q < len(A)-1; q++ {
			for r < len(A) && A[p]+A[q] > A[r] {
				r++
			}
			result = result + r - q - 1
		}
	}
	return result
}
