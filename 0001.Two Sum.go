func twoSum(nums []int, target int) []int {
	hashmap := make(map[int]int)
	var x int
	for i := 0; i < len(nums); i++ {
		x = target - nums[i]
		_, exist := hashmap[x]
		if exist {
			return []int{hashmap[x], i}
		} else {
			hashmap[nums[i]] = i
		}
	}
	panic("Something went wrong")
}

// 用map, nums = [2, 7, 11, 15] -> use a map[int]int to record
//                ^ -> search if target - nums[index] in map
// -> no : save in map[target - nums[index]] = index -> i++
//  nums[2, 7, 11, 15]
//          ^ -> if target - nums[index] = 2 in map
// -> yes : return [map[target - nums[index], index]

func twoSum(nums []int, target int) []int {
	var answer [2]int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				answer[0] = i
				answer[1] = j
				break
			}
		}
	}
	return answer[0:2]
}

// 暴力法： nums = [2, 7, 11, 15]
//                 ^  ^
//                 i. j.  -> if nums[i] + nums[j] == target -> return [i, j]
