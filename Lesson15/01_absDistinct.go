package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")
func getAbs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}
func Solution(A []int) int {
    // write your code in Go 1.4
    cnt := 0
    abs := make(map[int]bool)
    for _, a := range A {
        absA := getAbs(a)
        if abs[absA] == false {
            cnt = cnt + 1
            abs[absA] = true
        }
    } 
    return cnt
}