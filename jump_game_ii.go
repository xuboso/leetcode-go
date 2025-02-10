package main

import "fmt"

func main() {

	nums := []int{2, 3, 1, 1, 4}

	fmt.Println(jump(nums))
}

func jump(nums []int) int {
	curRight := 0  // 已建造的桥的右端点
	nextRight := 0 // 下一座桥的右端点的最大值
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
