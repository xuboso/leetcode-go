package main

import "fmt"

/**
 * 除自身以外数组的乘积
 * https://leetcode.cn/problems/product-of-array-except-self/
 */
func main() {

	nums := []int{1, 2, 3, 4}
	fmt.Println(ProductExceptSelf(nums))
}

func ProductExceptSelf(nums []int) []int {
	answer := make([]int, len(nums))
	product := 1
	for i := 0; i < len(nums); i++ {
		answer[i] = product
		product *= nums[i]
	}
	product = 1
	for i := len(nums) - 1; i >= 0; i-- {
		answer[i] *= product
		product *= nums[i]
	}
	return answer

}
