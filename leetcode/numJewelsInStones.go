package test

import (
	"fmt"
)

func main() {
	fmt.Println(numJewelsInStones("aA", "aAAbbbb"))
}
func numJewelsInStones(J string, S string) int {
	m := map[rune]bool{}
	for _, v := range J {
		m[v] = true
	}
	n := 0
	for _, v := range S {
		if _, ok := m[v]; ok {
			n += 1
		}
	}
	return n
}
