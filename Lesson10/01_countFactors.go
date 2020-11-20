package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(N int) int {
    // write your code in Go 1.4
    if N == 1 {
        return 1
    }
    result := 0
    curIndex := 1
    for {
        if curIndex * curIndex > N {
            break
        }
        if N % curIndex == 0{
            if (N / curIndex) == curIndex {
                result = result + 1
                break
            } else {
                result = result + 2
            }
        }
        curIndex = curIndex + 1
    }
    
    return result
}