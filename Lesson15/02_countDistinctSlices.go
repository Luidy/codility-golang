package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(M int, A []int) int {
	// write your code in Go 1.4
	p := 0
	q := 0
	result := 0
	distinctCheck := make([]int, M+1)

	for p < len(A) && q < len(A) {
		if distinctCheck[A[q]] == 0 {
			result = result + (q - p) + 1
			if result >= 1000000000 {
				return 1000000000
			}
			distinctCheck[A[q]] = 1
			q++
		} else {
			distinctCheck[A[p]] = 0
			p++
		}
	}
	return result
}
