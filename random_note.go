package main

import "fmt"

/**
 * 383. Ransom Note 赎金信
 * https://leetcode.com/problems/ransom-note/
 */
func main() {

	ransomoNote := "a"
	magazine := "b"
	fmt.Println(canConstruct(ransomoNote, magazine))
}

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	mp1 := map[byte]int{}
	for _, v := range magazine {
		mp1[byte(v)]++
	}

	fmt.Printf("mp1: %v\n", mp1)

	for _, k := range ransomNote {
		mp1[byte(k)]--
		if mp1[byte(k)] < 0 {
			return false
		}
	}

	return true
}
