package main

func canJump(nums []int) bool {
	max_j := 0
	if len(nums) <= 1 {
		return true
	}
	for i := 0; i < len(nums); i++ {
		if nums[i]+i > max_j+i {
			max_j = nums[i]
		}
		max_j--
		if max_j < 0 && i < len(nums)-1 {
			return false
		}
	}
	return true
}

// nums = [2, 3, 1, 1, 4] use a value max to record the current maximum jump length position
//         ^ i            max := 0.
//         if(nums[i] + i > max + i) -> max = nums[i]
//         2 + 0 > 0 -> max = nums[i] -> represent I can jump to index 2 -> max--, i++
// nums = [2, 3, 1, 1, 4] max = 1
//            ^ i         if(nums[i] + i > max + i) -> max = nums[i]
//                            3 + 1 = 4 > 1 + 1 -> max = nums[i], represent I can jump to index 4
// once max < 0 -> break -> return false -> else return true

// space complexity :O(1)
// time complexity :O(N)
