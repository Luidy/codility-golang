package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	N := len(A)
	mark := make([]int, N)

	for i := 0; i < N; i++ {
		mark[i] = -10000
	}

	mark[0] = A[0]
	for i := 0; i < N; i++ {
		for die := 1; die <= 6; die++ {
			next := i + die
			if next > N-1 {
				break
			}
			if mark[next] < mark[i]+A[next] {
				mark[next] = mark[i] + A[next]
			}
		}
	}
	return mark[N-1]
}
