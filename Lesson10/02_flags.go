package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	var peak []int
	for i := 1; i < len(A)-1; i++ {
		if A[i-1] < A[i] && A[i] > A[i+1] {
			peak = append(peak, i)
		}
	}
	if len(peak) <= 2 {
		return len(peak)
	}
	flag := 0
	for i := len(peak); i > 1; i-- {
		curFlag := 1
		curPos := peak[0]
		for x := 1; x < len(peak); x++ {
			if peak[x]-curPos >= i {
				curFlag++
				curPos = peak[x]
			} else {
				continue
			}
			if i == curFlag {
				break
			}
		}
		if flag < curFlag {
			flag = curFlag
		}
	}
	return flag
}
