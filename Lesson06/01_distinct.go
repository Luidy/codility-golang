package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
    d := make(map[int]bool)
    result := 0
    for _, a := range A {
        if _, v := d[a]; !v {
            d[a] = true
            result = result + 1
        }
    }
    
    return result
    // write your code in Go 1.4
}
