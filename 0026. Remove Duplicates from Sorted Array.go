package main

func removeDuplicates(nums []int) int {
	i := 0
	for j, _ := range nums {
		if nums[i] == nums[j] {
			j++
		} else {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

// nums = [0, 0, 1, 1, 1, 2, 2, 3, 3, 4] use two pointer i, j
//         ^i, j -> if nums[i] == nums[j] -> j start to search until nums[i] != nums[j]
// j++
// nums = [0, 0, 1, 1, 1, 2, 2, 3, 3, 4] use two pointer i, j
//         ^i ^j, j -> if nums[i] == nums[j] j++
// nums = [0, 0, 1, 1, 1, 2, 2, 3, 3, 4] use two pointer i, j
//         ^i    ^j, j -> if nums[i] != nums[j] i++
// nums = [0, 0, 1, 1, 1, 2, 2, 3, 3, 4] use two pointer i, j
//            ^i ^j, j -> nums[i] = nums[j], j++
// nums = [0, 1, 1, 1, 1, 2, 2, 3, 3, 4] use two pointer i, j
//            ^i    ^j -> if nums[i] == nums[j] -> j start to search until nums[i] != nums[j]
// ...
// return i + 1, i + 1 represents the remain number of "nums"
// j represents the index

// space complexity : O(1)
// time complexity : O(N)
