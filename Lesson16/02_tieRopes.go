package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(K int, A []int) int {
	// write your code in Go 1.4
	result := 0
	tie := 0
	for _, a := range A {
		tie = tie + a
		if tie >= K {
			result++
			tie = 0
		}
	}
	return result
}
