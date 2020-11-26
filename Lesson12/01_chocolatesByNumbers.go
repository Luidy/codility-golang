package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) []int {
    // write your code in Go 1.4
    countArr := make([]int, 2 * len(A)+1)
    result := make([]int, len(A))
    alreadyCal := make([]int, 2 * len(A)+1)
    
    for _, a := range A {
        countArr[a] = countArr[a] + 1
    }
    for i :=0; i<len(alreadyCal); i++ {
        alreadyCal[i] = -1
    }
    
    for i, a := range A {
        if alreadyCal[a] != -1 {
            result[i] = alreadyCal[a]
            continue
        }
        d := 1
        for {
            if d * d > a {
                break
            }
            if a % d == 0 {
                result[i] = result[i] + countArr[d]
                if d == (a / d) {
                    break
                } 
                result[i] = result[i] + countArr[a/d]
            }
            d = d + 1
        }
        result[i] = len(A) - result[i]
        alreadyCal[a] = result[i]
    }
    
    return result
}
