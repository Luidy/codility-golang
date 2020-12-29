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

func getMinIndex(start, end int, nail Matrix, index int) int {
	min := 0
	max := nail.Len() - 1
	minIndex := -1
	for min <= max {
		mid := (min + max) / 2
		if nail[mid][0] < start {
			min = mid + 1
		} else if nail[mid][0] > end {
			max = mid - 1
		} else {
			max = mid - 1
			minIndex = mid
		}
	}
	if minIndex == -1 {
		return -1
	}
	return minIndex
}

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
	fmt.Printf("%v \n", sortNail)
	result := 0
	for i := 0; i < len(A); i++ {
		result = getMinIndex(A[i], B[i], sortNail, result)
		fmt.Printf("A %d B %d i %d, minIndex %d \n", A[i], B[i], i, result)
	}
	return -1
}
