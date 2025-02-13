package main

import "fmt"

/**
 * 392. 判断子序列
 * https://leetcode.cn/problems/is-subsequence/
 */
func main() {

	s := "abc"
	t := "ahbgdc"
	fmt.Println(isSubsequence(s, t))
}

func isSubsequence(s string, t string) bool {
	n, m := len(s), len(t)
	i, j := 0, 0
	for i < n && j < m {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == n
}
