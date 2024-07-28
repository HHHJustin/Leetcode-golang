package main

func productExceptSelf(nums []int) []int {
	length := len(nums)
	res := make([]int, length)
	prefix := 1
	for i := 0; i < length; i++ {
		res[i] = prefix
		prefix *= nums[i]
	}
	postfix := 1
	for i := length - 1; i >= 0; i-- {
		res[i] *= postfix
		postfix *= nums[i]
	}
	return res
}

// nums   = [1, 2, 3, 4]
// output =  ^ except 1 -> 2 * 3 * 4 = 24
// nums   = [1, 2, 3, 4]
// output =     ^ except 2 -> 1 * 3 * 4 = 12
// nums   = [1, 2, 3, 4]
// output =        ^ except 1 -> 1 * 2 * 4 = 8
// ....

// and we can't use division operation and runs in O(n).

// prefix := 1
// res = [0, 0, 0, 0]
//        ^ = prefix and prefix *= nums[i] -> 1
// res = [1, 0, 0, 0] prefix = 1
//           ^ = prefix and prefix *= nums[i] -> 2
// res = [1, 1, 0, 0] prefix = 2
//              ^ = prefix and prefix *= nums[i] -> 6
// res = [1, 1, 2, 0]
//                 ^ = prefix -> 6
// res = [1, 1, 2, 6] and set postfix = 1

//.res = [1, 1, 2, 6]
//                 ^ = res[i] * postfix = 6 and postfix *= nums[i] = 4
//.res = [1, 1, 2, 6]
//              ^ = res[i] * postfix = 8 and postfix *= res[i] = 12
//.res = [1, 1, 8, 6]
//           ^ = res[i] * postfix = 12 and postfix *= res[i] = 24
//.res = [1, 12, 8, 6]
//        ^ = res[i] * postfix = 24
// res = [24, 12, 8, 6]

// res[0] -> use nothing in prefix step and use index 1, 2, 3 in postfix
// res[1] -> use index 0 in prefix step and use index 2, 3 in postfix
// res[2] -> use index 0, 1 in prefix step and use index 3 in postfix
// res[3] -> use index 0, 1, 2 in prefix step and use nothing in postfix

// space complexity : O(n)
// time complexity : O(n)
