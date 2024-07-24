package main

func maxProfit(prices []int) int {
	minindex := 0
	maxprofit := 0
	for index, price := range prices {
		maxprofit = max(maxprofit, price-prices[minindex])
		if price < prices[minindex] {
			minindex = index
		}
	}
	return maxprofit
}

// prices = [7, 1, 5, 3, 6, 4] -> if len(prices) <= 1 -> return 0
//              maxprofit := 0, minindex := 0
//              ^ compare prices[i] - prices[minindex] and maxprofit -> -6 and 0
//              maxprofit -> still
//              compare prices[i] and prices[minindex] -> minindex = i
// prices = [7, 1, 5, 3, 6, 4]
//             m^ i^ compare prices[i] - prices[minindex] and maxprofit -> 4 and 0
//              maxprofit = 4
//              compare prices[i] and prices[minindex] -> minindex still
// ...
// return maxprofit

// space complexity : O(1)
// time complexity : O(N)
