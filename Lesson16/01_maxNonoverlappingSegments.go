package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int, B []int) int {
    // write your code in Go 1.4
    if len(A) == 0 {
        return 0
    }
    lastEnd := B[0]
    result := 1
    for i, a := range A {
        if a > lastEnd {
            lastEnd = B[i]
            result = result + 1
        }
    }

    return result
}