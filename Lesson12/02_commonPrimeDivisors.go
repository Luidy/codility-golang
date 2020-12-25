package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")
func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
func Solution(A []int, B []int) int {
	// write your code in Go 1.4
	result := 0
	for i := 0; i < len(A); i++ {
		abGcd := gcd(A[i], B[i])
		a := A[i] / abGcd
		aGcd := gcd(a, abGcd)
		for aGcd != 1 {
			a = a / aGcd
			aGcd = gcd(aGcd, a)
		}
		b := B[i] / abGcd
		bGcd := gcd(b, abGcd)
		for bGcd != 1 {
			b = b / bGcd
			bGcd = gcd(bGcd, b)
		}
		if a == b {
			result++
		}
	}
	return result
}
