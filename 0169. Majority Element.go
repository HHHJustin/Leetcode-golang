package main

func majorityElement(nums []int) int {
	majority := nums[0]
	vote := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == majority {
			vote++
		} else {
			vote--
			if vote == 0 {
				majority = nums[i]
				vote = 1
			}
		}
	}
	return majority
}

// BM majority vote alorithm
// nums = [3, 2, 2]
// majority = nums[0], vote = 1
// nums = [3, 2, 2]
//            ^ -> if nums[i] != majority -> vote--
// nums = [3, 2, 2]
//            ^ -> if nums[i] != majority -> vote--
// if vote < 0 -> majority = nums[i], vote = 1
// ...
// return majority

// space complexity : O(1)
// time complexity : O(N)
