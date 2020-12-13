package solution

// you can also use imports, for example:
// import "fmt"
// import "os"
import "sort"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
    // write your code in Go 1.4
    if len(A) == 0 {
        return 1
    }
    sort.Ints(A)
    for i, a := range A {
        if a != i+1 {
            return i+1
        }
    }
    return len(A)+1
}