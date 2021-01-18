package solution

// you can also use imports, for example:
// import "fmt"
// import "os"
import "errors"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")
type Stack struct {
	nodes []int
	len   int
}

func (s *Stack) Push(n int) {
	s.nodes = append(s.nodes[:s.len], n)
	s.len++
}

func (s *Stack) Pop() (int, error) {
	if s.len == 0 {
		return -1, errors.New("No Data")
	}
	node := s.nodes[s.len-1]
	s.len--
	s.nodes = s.nodes[:s.len]
	return node, nil
}

func Solution(S string) int {
	// write your code in Go 1.4
	stack := Stack{}
	for i, _ := range S {
		if S[i:i+1] == "(" {
			stack.Push(1)
		} else if S[i:i+1] == ")" {
			_, err := stack.Pop()
			if err != nil {
				return 0
			}
		}
	}
	if stack.len == 0 {
		return 1
	}
	return 0

}
