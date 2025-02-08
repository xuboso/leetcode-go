package main

import (
	"fmt"
	"sort"
)

/**
 * https://leetcode.cn/problems/merge-sorted-array/?envType=study-plan-v2&envId=top-interview-150
 * 合并两个有序数组
 */
func main() {

	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3

	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2[:n])
	sort.Ints(nums1)
}
