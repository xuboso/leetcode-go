package main

import "fmt"

/**
 * 无重复字符的最长子串
 * https://leetcode.cn/problems/longest-substring-without-repeating-characters/
 * https://leetcode.cn/problems/longest-substring-without-repeating-characters/solutions/2883092/kao-yan-tvzhi-mian-dui-suo-zhao-hua-dong-xif8/?envType=study-plan-v2&envId=top-interview-150
 */
func main() {

	s := "abcabcbb"
	s = "abca"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	theHash := make(map[rune]int)
	result, left := 0, 0

	for right, letter := range s {
		if idx, found := theHash[letter]; found && idx >= left {
			left = idx + 1
		}
		theHash[letter] = right
		result = max(result, right-left+1)
		fmt.Printf("the hash: %v, left: %v, right: %v\n", theHash, left, right)

	}
	return result
}
