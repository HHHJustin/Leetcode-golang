package main

func jump(nums []int) int {
	max_l := 0
	total := 0
	end := 0
	for i := 0; i < len(nums)-1; i++ {
		max_l = max(max_l, nums[i]+i)
		if i == end {
			end = max_l
			total++
		}
	}
	return total
}

// nums = [2, 3, 1, 1, 4]
//.        ^ i -> use a value max_l to record the maximum length.           max_l := 0
//                use a value total to record the minimum number of jumps.  total := 0
//                use a value end to record current end of jump.            end := 0
// nums = [2, 3, 1, 1, 4]
//.           ^  ^ mean I can jump to index 1 ~ 2 -> use end to record -> end = 2,
//                 and mean I have two choice for jumping 1 or 2 -> total++,
//                 compare max_l & num[i] + i in range of index 1 ~ 2
// ...
// nums = [2, 3, 1, 1, 4] so once i == end to update the end and total++
//.              ^ -> end = max_l nums[1] ~ nums[2] can arrive maximum jump.
// ...
// return total

// space compplexity : O(1)
// time complexity : O(N)
