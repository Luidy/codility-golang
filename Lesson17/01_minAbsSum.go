package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")
func abs(a int) int {
    if a >= 0 {
        return a
    }
    return -a
}

func Solution(A []int) int {
    // write your code in Go 1.4
    maxLen := 0
    countNum := make([]int, 101)
    for _, a := range A {
        absNum := abs(a)
        countNum[absNum] = countNum[absNum] + 1
        maxLen = maxLen + absNum
    }

    subset := make([]int, maxLen/2+1)
    subset[0] = 1
    for x, v := range countNum {
        if v == 0 {
            continue
        }
        for i:=0; i<maxLen/2+1; i++ {
            if subset[i] > 0 {
                subset[i] = v+1
            } else if (i >= x && subset[i-x] > 1) {
                subset[i] = subset[i-x] - 1
            }
        }
    }

    result := maxLen
    for k, s := range subset {
        if s == 0 {
            continue
        }
        if result > maxLen-k*2 {
            result =  maxLen-k*2
        }
    }

    return result
}