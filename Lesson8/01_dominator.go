package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")
func Solution(A []int) int {
    // write your code in Go 1.4
    persons := make(map[int]int)
    maxCount := len(A)/2
    dominator := -1
    for i, a := range A {
        persons[a] = persons[a] + 1
        if persons[a] > maxCount {
            maxCount = persons[a]
            dominator = i
            print(a)
        }
        
    }

    return dominator
    
}