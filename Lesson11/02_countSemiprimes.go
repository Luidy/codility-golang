package solution

// you can also use imports, for example:
import "fmt"

// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(N int, P []int, Q []int) []int {
	// write your code in Go 1.4
	prime := make([]int, N+1)
	i := 2
	for i*i < N {
		if prime[i] == 0 {
			k := i * i
			for k <= N {
				if prime[k] == 0 {
					prime[k] = i
				}
				k = k + i
			}
		}
		i++
	}
	fmt.Printf("%v", prime)
	result := make([]int, len(P))
	return result
}
