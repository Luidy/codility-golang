package solution

// you can also use imports, for example:
// import "fmt"
// import "os"
import "math"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")
func fibo(len, mod int) []int {
	fibo := make([]int, len+1)
	for i, _ := range fibo {
		if i == 0 || i == 1 {
			fibo[i] = 1
			continue
		}
		fibo[i] = (fibo[i-1] + fibo[i-2]) % mod
	}
	return fibo
}
func Solution(A []int, B []int) []int {
	// write your code in Go 1.4
	maxFibo := 0
	for _, a := range A {
		if a > maxFibo {
			maxFibo = a
		}
	}
	maxMod := 0
	for _, b := range B {
		if b > maxMod {
			maxMod = b
		}
	}
	maxMod = int(math.Pow(2, float64(maxMod)))
	fibo := fibo(maxFibo, maxMod)
	result := make([]int, len(A))
	for i := 0; i < len(A); i++ {
		result[i] = fibo[A[i]] % int(math.Pow(2, float64(B[i])))
	}
	return result
}
