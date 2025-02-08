package main

import "fmt"

/**
 * 移除元素
 * https://leetcode.cn/problems/remove-element/
 */
func main() {

	nums := []int{3, 2, 2, 3}
	val := 3
	different_num := removeElement(nums, val)
	fmt.Printf("nums : %v, different num is : %v\n", nums, different_num)
}

func removeElement(nums []int, val int) int {
	left := 0
	for _, v := range nums {
		if v != val {
			nums[left] = v
			left++
		}
	}
	return left
}
