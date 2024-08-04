package main

import "fmt"

func summaryRanges(nums []int) []string {
	output := new([]string)
	if len(nums) == 0 {
		return *output
	}
	left, right, temp := 0, 0, nums[0]
	var s string
	for right < len(nums) {

		if nums[right] == nums[left] {
			right++
		} else if nums[right] == temp+1 {
			temp = nums[right]
			right++
		} else if nums[right] != temp+1 {
			if nums[left] == nums[right-1] {
				s = fmt.Sprintf("%d", nums[left])
			} else {
				s = fmt.Sprintf("%d->%d", nums[left], nums[right-1])
			}
			*output = append(*output, s)
			temp = nums[right]
			left = right
		}
	}
	if nums[left] == nums[len(nums)-1] {
		s = fmt.Sprintf("%d", nums[left])
	} else {
		s = fmt.Sprintf("%d->%d", nums[left], nums[right-1])
	}
	*output = append(*output, s)
	return *output
}

// make a string slice output and use 2 pointer
// nums = [0, 1, 2, 4, 5, 7]
//         ^ left, right     -> right to check next if == nums[left] + 1
//                           -> if not -> make a string fmt.sprintf("%d->%d",nums[left],nums[righy])
//                           -> and append(output, string)
// return output

// space complexity: O(N)
// time complexity: O(N)
