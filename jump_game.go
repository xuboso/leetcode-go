package main

import "fmt"

func main() {

	nums := []int{2, 3, 1, 1, 4}

	fmt.Println(canJump(nums))
}

func canJump(nums []int) bool {
	mx := 0
	for i, jump := range nums {
		if i > mx {
			return false
		}
		mx = max(mx, i+jump)
	}
	return true
}
