package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(X int, A []int) int {
    // write your code in Go 1.4
    if len(A) < X {
        return -1
    }
    
    d := make([]bool, len(A))
    p := 0
    for i, a := range A {
        if d[a-1] == false && a<=X {
            d[a-1] = true
            p = p + 1
            if p == X {
                return i
            }
        } 
    }
    return -1
}