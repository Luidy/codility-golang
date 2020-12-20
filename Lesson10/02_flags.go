package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	peak := make([]int, len(A))
	peakCnt := 0
	for i := 1; i < len(A)-1; i++ {
		if A[i-1] < A[i] && A[i] > A[i+1] {
			peak[i] = 1
			peakCnt++
		}
	}
	if peakCnt <= 2 {
		return peakCnt
	}
	next := make([]int, len(A))
	for i := len(A) - 2; i >= 0; i-- {
		if peak[i] == 1 {
			next[i] = i
		} else {
			next[i] = next[i+1]
		}
	}

	flag := 0
	i := 1
	for (i-1)*i <= len(A) {
		curFlag := 0
		curPos := 0
		for curFlag < i && curPos < len(A) {
			curPos = next[curPos]
			if curPos == 0 {
				break
			}
			curFlag++
			curPos = curPos + i
		}

		if flag < curFlag {
			flag = curFlag
		}
		i++
	}
	return flag
}
