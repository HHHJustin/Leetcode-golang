package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := len(nums1) - 1; i >= 0; i-- {
		if n-1 < 0 || (m-1 >= 0 && nums1[m-1] >= nums2[n-1]) {
			// n - 1要放在前面,如果先做後面會報out of index.
			nums1[i] = nums1[m-1]
			m--
		} else {
			nums1[i] = nums2[n-1]
			n--
		}
	}
}

// nums1 = [1, 2, 3, 0 ,0 ,0] nums2 = [2, 5, 6]
//                ^ m-1    ^ i               ^ n-1
// i := len(nums1) - 1
// nums1 = [1, 2, 3, 0 ,0 ,6] nums2 = [2, 5, 6]
//                ^ m-1 ^ i               ^ n-1
// nums1 = [1, 2, 3, 0 ,5 ,6] nums2 = [2, 5, 6]
//                ^  ^ i               ^ n-1
// nums1 = [1, 2, 3, 3 ,5 ,6] nums2 = [2, 5, 6]
//             ^  ^ i.                 ^ n-1
// .....

// func merge(nums1 []int, m int, nums2 []int, n int)  {
//     i := 0
//     j := 0
//     for m > 0 && n > 0 {
//         if nums1[i] <= nums2[j] {
//             m--
//             i++
//         } else {
//             for k:= len(nums1) - 1; k > i ; k-- {
//                 nums1[k] = nums1[k - 1]
//             }
//             nums1[i] = nums2[j]
//             fmt.Println(nums1)
//             n--
//             i++
//             j++
//         }
//     }
//     if m == 0 {
//         for n > 0 {
//             nums1[i] = nums2[j]
//             i++
//             j++
//             n--
//         }
//     }
// }

// assume len_nums1 and len_nums2 to record the remain number
// num1 = [1, 2, 3, 0, 0, 0]  nums2 = [2, 5, 6]
//.        ^ i                         ^ j
// -> if nums1[i] <= nums2[j] -> i++ len_nums1 --
// num1 = [1, 2, 3, 0, 0, 0]  nums2 = [2, 5, 6]
//            ^ i                      ^ j
// -> if nums1[i] <= nums2[j] -> i++ len_nums1 --
// num1 = [1, 2, 3, 0, 0, 0]  nums2 = [2, 5, 6]
//               ^ i                   ^ j
// -> if nums1[i] > nums2[j] -> num1 往後移一格 -> nums1[i] = nums2[j] j++ len_nums2 --
// num1 = [1, 2, 2, 3, 0, 0]  nums2 = [2, 5, 6]
//                  ^ i                   ^ j
// -> if nums1[i] <= nums2[j] -> i++ len_nums1 --
// -> len_nums1 == 0 將剩餘nums2後面都填到nums1後
