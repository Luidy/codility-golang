package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")
type Stack struct {
	items []int
	size  int
}

func (s *Stack) Pop() int {
	if s.size == 0 {
		return -1
	}
	result := s.items[s.size-1]
	s.items = s.items[0 : s.size-1]
	s.size = s.size - 1
	return result
}

func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
	s.size = s.size + 1
}

func Solution(A []int, B []int) int {
	// write your code in Go 1.4
	s := Stack{}
	passFish := 0
	for i, b := range B {
		if b == 1 {
			s.Push(A[i])
		} else {
			for s.size > 0 {
				curFish := s.Pop()
				if A[i] > curFish {
					continue
				} else {
					s.Push(curFish)
					break
				}
			}
			if s.size == 0 {
				passFish = passFish + 1
			}
		}
	}
	return s.size + passFish
}
