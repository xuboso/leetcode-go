package main

import (
	"fmt"
	"sort"
)

/**
 * 多数元素
 * https://leetcode.cn/problems/majority-element/
 */
func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}

	fmt.Println(majorityElement(nums))
}

func majorityElement(nums []int) int {

	len := len(nums)
	sort.Ints(nums)
	return nums[len/2]
}
