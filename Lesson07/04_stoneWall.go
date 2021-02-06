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

func (s *Stack) Top() int {
	if s.size == 0 {
		return -1
	}
	result := s.items[s.size-1]
	return result
}

func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
	s.size = s.size + 1
}

func Solution(H []int) int {
	// write your code in Go 1.4
	stack := Stack{}
	wall := 0
	for i := 0; i < len(H); i++ {
		item := stack.Top()
		prev := H[i]
		if item != -1 && item > H[i] {
			for item > H[i] {
				if prev != item {
					wall++
				}
				prev = stack.Pop()
				item = stack.Top()
			}
		}
		stack.Push(H[i])
	}
	cur := -1
	top := stack.Pop()
	for top != -1 {
		if cur != top {
			wall++
		}
		cur = top
		top = stack.Pop()
	}

	return wall
}
