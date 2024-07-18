package main

func removeElement(nums []int, val int) int {
	i := 0
	j := 0
	for j < len(nums) {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
			j++
		} else {
			j++
		}
	}
	return i
}

// for -> j := 0 ; j < len(nums)
// nums = [0, 1, 2, 2, 3, 0, 4, 2] , val = 2
//         ^i, j -> if nums[j] != val -> nums[i] = nums[j], i++, j++
// nums = [0, 1, 2, 2, 3, 0, 4, 2] , val = 2
//            ^i, j -> if nums[j] != val -> nums[i] = nums[j], i++, j++
// nums = [0, 1, 2, 2, 3, 0, 4, 2] , val = 2
//               ^i, j -> if nums[i] == val -> use j to search until nums[j] != val
// nums = [0, 1, 2, 2, 3, 0, 4, 2] , val = 2
//               ^i.   ^j then nums[i] = nums[j], i++, j++
// nums = [0, 1, 3, 2, 3, 0, 4, 2] , val = 2
//                  ^i.   ^j. -> if nums[j] != val -> nums[i] = nums[j], i++, j++
// ...
// return i -> i represents the number of elements in the array
// space complexity : O(1)
// time complexity : O(N)
