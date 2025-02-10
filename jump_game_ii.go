package main

import "fmt"

func main() {

	nums := []int{2, 3, 1, 1, 4}

	fmt.Println(jump(nums))
}

func jump(nums []int) int {
	curRight := 0
	nextRight := 0
	count := 0

	for i, num := range nums[:len(nums)-1] {
		nextRight = max(nextRight, i+num)
		if i == curRight {
			curRight = nextRight
			count++
		}
	}

	return count
}
