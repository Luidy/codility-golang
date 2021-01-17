package solution

// you can also use imports, for example:
// import "fmt"
// import "os"
import "sort"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	lower := make([]int, len(A))
	upper := make([]int, len(A))
	for i, a := range A {
		lower[i] = i - a
		upper[i] = i + a
	}
	sort.Ints(lower)
	sort.Ints(upper)

	intersect := 0
	l := 0
	for u := 0; u < len(A); u++ {
		for l < len(A) && upper[u] >= lower[l] {
			intersect += l
			intersect -= u
			l++
		}
	}
	if intersect > 10000000 {
		return -1
	}
	return intersect
}
