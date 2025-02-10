package main

import "fmt"

/**
 * 135. 分发糖果
 * https://leetcode.cn/problems/candy/
 * https://leetcode.cn/problems/candy/solutions/2777041/xiao-zhou-ti-jie-135-fen-fa-tang-guo-by-ayt3x
 */
func main() {
	ratings := []int{1, 2, 2}
	fmt.Println(candy(ratings))
}

func candy(ratings []int) int {
	n := len(ratings)

	left := make([]int, n)
	left[0] = 1
	for i := 1; i <= n-1; i++ {
		if ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}

	right := make([]int, n)
	right[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			right[i] = right[i+1] + 1
		} else {
			right[i] = 1
		}
	}

	res := 0
	for i := 0; i < n; i++ {
		res += max(left[i], right[i])
	}

	return res
}
