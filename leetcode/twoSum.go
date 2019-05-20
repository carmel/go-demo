package test

import (
	"fmt"
)

func main() {
	arr := []int{3, 2, 4}
	fmt.Println(twoSum(arr, 6))
}

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	l := len(nums)

	for i := 0; i < l; i++ {
		m[nums[i]] = i
	}
	fmt.Println(m)
	var s int
	for i := 0; i < l; i++ {
		s = target - nums[i]
		fmt.Println(i, nums[i], s)
		if k, ok := m[s]; ok && k != i {
			return []int{i, m[s]}
		}
	}
	return nil
}
