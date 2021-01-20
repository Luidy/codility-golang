package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(N int) int {
	// write your code in Go 1.4
	sum := N + 1
	i := 2
	for i*i <= N {
		if N%i == 0 {
			curSum := i + N/i
			sum = curSum
		}
		i++
	}
	return sum * 2
}
