package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")
func Max(a, b int) int {
    if a < b {
        return b
    }
    
    return a
}

func Solution(A []int) int {
    // write your code in Go 1.4
    fSum := make([]int, len(A))
    bSum := make([]int, len(A))
    fSum[0] = 0
    bSum[len(bSum)-1] = 0
    
    for i := 1; i<len(A); i++ {
        fSum[i] = Max(fSum[i-1]+A[i], 0)
        bSum[len(bSum)-1-i] = Max(bSum[len(bSum)-i]+A[len(bSum)-1-i], 0)
    }
    
    max := 0
    for i := 1; i<len(A)-1; i++ {
        curSum := fSum[i-1]+bSum[i+1]
        if curSum > max {
            max = curSum
        }
    }
    
    return max
}