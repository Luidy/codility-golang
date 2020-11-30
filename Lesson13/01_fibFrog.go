package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")
func fibo(N int) (fibArr []int) {
    fibArr = append(fibArr, 0)
    fibArr = append(fibArr, 1)
    fibArr = append(fibArr, 1)

    i:=2
    for fibArr[i] < N {
        i = i+1
        fibArr = append(fibArr, fibArr[i-1] + fibArr[i-2])
    }

    return fibArr
}

func Solution(A []int) int {
    // write your code in Go 1.4
    A = append(A, 1)
    N := len(A)
    fibArr := fibo(N)
    if fibArr[len(fibArr)-1] == N {
        return 1
    } 
    fibArr = fibArr[1:len(fibArr)-1]
    
    route := make([]int, N)
    for _, f := range fibArr {
        if A[f-1] == 1 {
            route[f-1] = 1
        }
    }

    for i:=0; i<N; i++{
        if A[i] != 1 || route[i] > 0{
            continue
        }

        minJump := i+1
        canJump := false
        for _, f := range fibArr {
            if i-f < 0 || route[i-f] == 0 {
                continue;
            }
            if minJump > route[i-f] {
                minJump = route[i-f]
                canJump = true
            }
        }
        if canJump {
            route[i] = minJump + 1
        }
    }
    if route[len(route)-1] == 0 {
        return -1
    }
    return route[len(route)-1]
}