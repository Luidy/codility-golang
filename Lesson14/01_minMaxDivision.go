package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")
func isDivisible(mid, K int, A []int) bool {
    sum := 0
    for _, a := range A {
        sum = sum + a
        if sum > mid {
            K = K - 1
            sum = a
            if K == 0 {
                return false
            }
        }
    }
    return true
}

func Solution(K int, M int, A []int) int {
    // write your code in Go 1.4
    beg := 0
    end := 0
    for _, a := range A {
        if a > beg {
            beg = a
        }
        end = end + a
    }

    min := end
    for (beg <= end) {
        mid := (beg+end) / 2
        if isDivisible(mid, K, A) {
            min = mid
            end = mid - 1
        } else {
            beg = mid + 1
        }
    }
    return min

}