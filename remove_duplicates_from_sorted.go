package main

import "fmt"

/**
 * 删除排序数组中的重复项
 * https://leetcode.cn/problems/remove-duplicates-from-sorted-array/
 */
func main() {

	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	different_num := removeDuplicates(nums)
	fmt.Printf("different num is : %v\n", different_num)
}

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	slow := 1
	for fast := 1; fast < n; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
