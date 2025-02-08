package main

import "fmt"

/**
 * 删除排序数组中的重复项 II
 * https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii/
 */
func main() {

	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	len := removeDuplicatesII(nums)
	fmt.Println("len : ", len)
}

func removeDuplicatesII(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}

	slow, fast := 2, 2
	for fast < n {
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
		fmt.Println(nums)
	}
	return slow
}
