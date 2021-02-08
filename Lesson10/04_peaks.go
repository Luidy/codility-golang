package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	N := len(A)
	peaks := []int{}
	for i := 1; i < N-1; i++ {
		if A[i] > A[i-1] && A[i] > A[i+1] {
			peaks = append(peaks, i)
		}
	}

	if len(peaks) == 0 {
		return 0
	}

	for i := len(peaks); i > 1; i-- {
		if N%i != 0 {
			continue
		}

		cur := 0
		k := N / i
		for _, p := range peaks {
			if cur*k <= p && (cur+1)*k > p {
				cur++
			}
		}
		if cur == i {
			return i
		}
	}
	return 1
}
