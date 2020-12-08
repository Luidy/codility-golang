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
    countNum := make(map[int]int)
    for _, a := range A {
        absNum := abs(a)
        countNum[absNum] = countNum[absNum] + 1
        maxLen = maxLen + absNum
    }

    subset := make(map[int]int)
    subset[0] = 1
    for k, v := range countNum {
        for i:=0; i<maxLen/2+1; i++ {
            if subset[i] > 0 {
                subset[i] = v+1
            } else if (i >= k && subset[i-k] > 1) {
                subset[i] = subset[i-k] - 1
            }
        }
    }

    result := maxLen
    for k, _ := range subset {
        if result > maxLen-k*2 {
            result =  maxLen-k*2
        }
    }

    return result
}