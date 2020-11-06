package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int, K int) []int {
    if len(A) == 0 {
        return A
    }
    
    result := make([]int, len(A))
    for i, a := range A {
        result[(i+K)%len(A)] = a
    }
    print(result[0])
    return result
}