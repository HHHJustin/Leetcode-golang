package main

func swap(slice []int, i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func reverse(slice []int, start, last int) {
	for start < last {
		swap(slice, start, last)
		start++
		last--
	}
}

func rotate(nums []int, k int) {
	if len(nums) <= 1 {
		return
	}
	k %= len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

// *k need mod len(nums)
// nums = [1, 2, 3, 4, 5, 6, 7], k = 3
// -> mean to rotate 3 steps
// -> rotate 1 steps to right = [7, 1, 2, 3, 4, 5, 6]
// -> rotate 2 steps to right = [6, 7, 1, 2, 3, 4, 5]
// -> rotate 3 steps to right = [5, 6, 7, 1, 2, 3, 4]
// Step
// 1. reverse all array -> [7, 6, 5, 4, 3, 2, 1]
// 2. reverse array[0 : k] -> [5, 6, 7, 4, 3, 2, 1]
// 3. reverse array[k : len(nums)] -> [5, 6, 7, 1, 2, 3, 4]
// so we need a function reverse

// reverse pseudo code
// ex: num = [1, 2, 3, 4, 5]
//            ^i          ^j -> swap(i, j), i++, j-- -> i == 1, j == 3, if i < j -> continue
// num = [5, 2, 3, 4, 1]
//           ^i    ^j -> swap(i, j), i++, j-- -> i == 2, j == 2, if i >= j -> end.

// space complexity : O(1)
// time complexity : O(N)
