package main

import "fmt"

/**
 * 罗马数字转整数
 * https://leetcode.cn/problems/roman-to-integer/
 */
func main() {
	s := "MCMXCIV"
	fmt.Println(romanToInt(s))
}

func romanToInt(s string) int {
	symbolValues := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	n := len(s)
	ans := 0

	for i := range s {
		value := symbolValues[s[i]]
		if i < n-1 && value < symbolValues[s[i+1]] {
			ans -= value
		} else {
			ans += value
		}
	}

	return ans
}
