package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(N int, A []int) []int {
    // write your code in Go 1.4
    tempMax, max := 0, 0
    result := make([]int, N)
    for _, a := range A {
        if a == N + 1 {
            max = tempMax
        } else {
            if result[a-1] < max {
                result[a-1] = max + 1
            } else {
                result[a-1]++
            }
            if tempMax < result[a-1] {
                tempMax = result[a-1]
            }
        }
    }

    for i, r := range result {
        if r < max {
            result[i] = max
        }
    }

    return result 
}