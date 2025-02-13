package main

import "fmt"

/**
 * 两数之和
 * https://leetcode.cn/problems/two-sum/
 */
func main() {

	nums := []int{3, 2, 4}
	fmt.Println(twoSum(nums, 6))
}

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for i, num := range nums {
		if v, ok := hashMap[target-num]; ok {
			return []int{v, i}
		}
		hashMap[num] = i
	}
	return nil
}
