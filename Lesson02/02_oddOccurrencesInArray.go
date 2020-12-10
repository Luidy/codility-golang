package solution

// you can also use imports, for example:
// import "fmt"
// import "os"
import "sort"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
    // write your code in Go 1.4
    if len(A) == 1 {
        return A[0]
    }
    sort.Ints(A)
    for i:=0; i<len(A)-1; i = i+2 {
        if A[i] != A[i+1] {
            return A[i]
        }
    }
    return A[len(A)-1]
}