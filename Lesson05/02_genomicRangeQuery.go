package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")
func if2int(importFactor string)int{
    if importFactor == "A" {
        return 1
    } else if importFactor == "C" {
        return 2
    } else if importFactor == "G" {
        return 3
    }else if importFactor == "T" {
        return 4
    }
    return 0
}

func Solution(S string, P []int, Q []int) []int {
    // write your code in Go 1.4
    ifCnt := make([][4]int, len(S)+1)
    var curIfCnt = [4]int{0, 0, 0, 0}

    for i, s := range S {
        ifInt := if2int(string(s))
        curIfCnt[ifInt - 1] = curIfCnt[ifInt - 1] + 1
        ifCnt[i+1] = curIfCnt
    }
    result := make([]int, len(P))
    for i:=0; i<len(P); i++ {
        for j:=0; j<4; j++ {
            if ifCnt[Q[i]+1][j] - ifCnt[P[i]][j] != 0 {
                result[i] = j+1
                break
            }
        }
    }
    return result
}