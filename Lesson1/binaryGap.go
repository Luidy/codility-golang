package solution

// you can also use imports, for example:
import (
    "fmt"
    "strings"
)
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(N int) int {
    b := fmt.Sprintf("%b", N)
    if !strings.Contains(b, "0") || strings.Count(b, "1") == 1{
        return 0
    }
    
    cur :=0
    max :=0
    for _, v := range string(b){
        if string(v) == "1"{
            if cur > max{
                max = cur
            }
            cur = 0
        } else {
            cur = cur + 1
        }
    }
    
    return max
    // write your code in Go 1.4
}