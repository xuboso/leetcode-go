package main

import "fmt"

/**
 * 最后一个单词的长度
 * https://leetcode.cn/problems/length-of-last-word/
 */
func main() {
	word := "   fly me   to   the moon  "
	fmt.Println(lengthOfLastWord(word))
}

func lengthOfLastWord(s string) int {
	index := len(s) - 1
	ans := 0
	for s[index] == ' ' {
		index--
	}

	for index >= 0 && s[index] != ' ' {
		ans++
		index--
	}
	return ans
}
