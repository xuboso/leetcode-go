package main

/**
 * 买卖股票的最佳时机
 * https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/
 */
func main() {

	prices := []int{7, 1, 5, 3, 6, 4}
	ans := maxProfit(prices)
	println(ans)
}

func maxProfit(prices []int) (ans int) {
	minPrice := prices[0]
	for _, p := range prices {
		minPrice = min(minPrice, p)
		ans = max(ans, p-minPrice)
	}
	return
}
