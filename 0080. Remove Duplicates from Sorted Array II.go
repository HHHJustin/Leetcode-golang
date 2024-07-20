package main

func removeDuplicates2(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	count := 0
	i := 2
	if nums[0] == nums[1] {
		count = 1
	}
	for j, _ := range nums {
		if j >= 2 {
			if nums[j-1] == nums[j] {
				if count == 0 {
					nums[i] = nums[j]
					i++
					count++
				}
			} else {
				nums[i] = nums[j]
				i++
				count = 0
			}
		}

	}
	return i
}

// nums = [0, 0, 1, 1, 1, 1, 2, 3, 3]  use two pointers and a value count record appear times
// check if len(nums) <= 1 -> return nums
// check if nums[0] == nums[1] -> yes -> count = 1
// search nums[2:]
// nums = [0, 0, 1, 1, 1, 1, 2, 3, 3]
//               ^i, j -> nums[j] != nums[j - 1] -> count = 0, i++, j++
// nums = [0, 0, 1, 1, 1, 1, 2, 3, 3]
//                  ^i, j -> nums[j] == nums[j - 1] && count = 0 -> count = 1, i++, j++
// nums = [0, 0, 1, 1, 1, 1, 2, 3, 3]
//                     ^i, j -> nums[j] == nums[j - 1] && count = 1 -> j++
// nums = [0, 0, 1, 1, 1, 1, 2, 3, 3]
//                     ^i ^ j -> nums[j] == nums[j - 1] && count = 1 -> j++
// nums = [0, 0, 1, 1, 1, 1, 2, 3, 3]
//                     ^i    ^ j -> nums[j] != nums[j - 1] -> nums[i] = nums[j], count = 0, i++, j++
// nums = [0, 0, 1, 1, 2, 1, 2, 3, 3]
//                        ^i    ^ j -> nums[j] != nums[j - 1] -> nums[i] = nums[j], count = 0, i++, j++
// nums = [0, 0, 1, 1, 2, 3, 2, 3, 3]
//                           ^i    ^ j -> nums[j] != nums[j - 1]
// -> nums[i] = nums[j], count = 0, i++, j++
// nums = [0, 0, 1, 1, 2, 3, 2, 3, 3]
//                           ^i    ^ j -> nums[j] == nums[j - 1] && count = 0
// -> nums[i] = nums[j], count++, i++, j++
// return i + 1
