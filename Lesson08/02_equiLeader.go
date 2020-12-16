package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	a_map := make(map[int]int)
	leader := 0
	leaderCnt := 0
	for _, a := range A {
		a_map[a] = a_map[a] + 1
		if a_map[a] > leaderCnt {
			leaderCnt = a_map[a]
			leader = a
		}
	}

	leftLeader := 0
	rightLeader := a_map[leader]
	result := 0
	for i, a := range A {
		if a == leader {
			leftLeader = leftLeader + 1
			rightLeader = rightLeader - 1
		}
		if (i+1)/2 < leftLeader && (len(A)-i-1)/2 < rightLeader {
			result = result + 1
		}
	}
	return result
}
