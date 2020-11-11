package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(S string) int {
    // write your code in Go 1.4
    if S == "" {
        return 1
    }
    pair := map[string]string{
        ")" : "(",
        "}" : "{",
        "]" : "[",
        "V" : "OK",
        "W" : "OK",
        "U" : "OK",
        "T" : "OK",
    }
    
    stack := []string{}
    
    for _, item := range S {
        s := string(item)
        if pair[s] == "OK" {
            continue
        } else if len(stack) == 0 {
            stack = append(stack, s)
            continue
        } else if pair[s] == stack[len(stack)-1] {
            stack = stack[:len(stack)-1]
            continue
        } else {
            stack = append(stack, s)
            continue
        }
    }
    if len(stack) == 0 {
        return 1
    }
    return 0
}
