package main

import "fmt"

/**
 * 最长公共前缀
 * https://leetcode.cn/problems/longest-common-prefix/
 */
func main() {

	strs := []string{"flower", "flow", "flight"}

	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	s0 := strs[0]

	for j := range s0 {
		for _, s := range strs {
			if j == len(s) || s[j] != s0[j] {
				return s0[:j]
			}
		}
	}

	return s0
}
