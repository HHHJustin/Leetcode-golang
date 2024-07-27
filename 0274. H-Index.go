package main

import (
	"sort"
)

func hIndex(citations []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(citations)))
	var i int
	for i = 0; i < len(citations); i++ {
		if citations[i] < i+1 {
			break
		}
	}
	return i
}

// citations = [3, 0, 6, 1, 5] -> descending sorting
// citations = [6, 5, 3, 1, 0]
//              ^ i, if citations[i] >= i + 1, next
// citations = [6, 5, 3, 1, 0]
//                 ^ i, if citations[i] >= i + 1, next
// citations = [6, 5, 3, 1, 0]
//                    ^ i, if citations[i] >= i + 1, next
// citations = [6, 5, 3, 1, 0]
//                       ^ i, if citations[i] < i + 1, return i.

// space complexity : O(1)
// time complexity : O(N)
