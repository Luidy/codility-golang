package solution

// you can also use imports, for example:
import (
	"fmt"
	"sort"
)

// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")
type Matrix [][]int

func (m Matrix) Len() int           { return len(m) }
func (m Matrix) Less(i, j int) bool { return m[i][0] < m[j][0] }
func (m Matrix) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

func Solution(A []int, B []int, C []int) int {
	// write your code in Go 1.4
	sortNail := make(Matrix, len(C))
	for i := 0; i < len(C); i++ {
		sortNail[i] = make([]int, 2)
	}
	for i, c := range C {
		sortNail[i][0] = c
		sortNail[i][1] = i
	}
	sort.Sort(sortNail)
	fmt.Printf("%v", sortNail)
	return -1
}
