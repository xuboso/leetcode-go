package main

import "fmt"

func main() {
	word := "   fly me   to   the moon  "
	fmt.Println(LastWord(word))
}

func LastWord(s string) string {
	index := len(s) - 1
	for s[index] == ' ' {
		index--
	}

	start := index
	for start >= 0 && s[start] != ' ' {
		start--
	}

	return s[start+1 : index+1]
}
