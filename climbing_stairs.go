package main

import "fmt"

/**
 * 爬楼梯
 * https://leetcode.cn/problems/climbing-stairs/
 */
func main() {

	total := climbStairs(45)
	fmt.Println("递归 : ", total)

	total = climbStairs2(45)
	fmt.Println("动态规划 : ", total)
}

func climbStairs(n int) int {

	if n < 2 {
		return 1
	}

	if n == 2 {
		return 2
	}

	return climbStairs(n-1) + climbStairs(n-2)
}

func climbStairs2(n int) int {

	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
