package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := len(nums1) - 1; i >= 0; i-- {
		if n-1 < 0 || (m-1 >= 0 && nums1[m-1] >= nums2[n-1]) {
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
// while i > 0, i-- each time
// if nums1[m-1] >= nums2[n-1] || m - 1 < 0 -> nums1[i] = nums1[m-1], m--
// else -> nums1[i] = nums2[n-1], n--

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
