package main

func maxProfit2(prices []int) int {
	min := 0
	sum := 0
	maxprofit := 0
	if len(prices) <= 1 {
		return 0
	}
	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1] < 0 {
			if maxprofit > 0 {
				sum += maxprofit
				maxprofit = 0
			}
			min = i
		} else {
			maxprofit = prices[i] - prices[min]
		}
	}
	if maxprofit > 0 {
		sum += maxprofit
	}
	return sum
}

// prices = [7, 1, 5, 3, 6, 4] use value min to record the minimum index
//           ^                 use value sum to record total profit
//                             use maxprofit to record temp  profit
// prices = [7, 1, 5, 3, 6, 4]
//       min ^  ^ i -> check if prices[i] - prices[i - 1] < 0
//                  -> yes -> min = i, i++
// prices = [7, 1, 5, 3, 6, 4]
//          min ^  ^ i -> check if prices[i] - prices[i - 1] < 0
//                     -> no -> maxprofit = prices[i] - prices[min]
// prices = [7, 1, 5, 3, 6, 4]
//          min ^  ^ i -> check if prices[i] - prices[i - 1] < 0
//                     -> yes && maxprofit != 0 -> sum += maxprofit, maxprofit = 0, min = i, i++
// ...
// return sum

// space complexity : O(1)
// time complexity : O(N)
