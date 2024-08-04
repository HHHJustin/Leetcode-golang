package main

func canCompleteCircuit(gas []int, cost []int) int {
	total := 0
	temp := 0
	index := 0
	for i := 0; i < len(gas); i++ {
		diff := gas[i] - cost[i]
		temp += diff
		if temp < 0 {
			temp = 0
			index = (i + 1) % len(gas)
		}
		total += diff
	}
	if total < 0 {
		return -1
	} else {
		return index
	}
}

// gas    = [1,   2, 3,  4, 5]
// cost   = [3,   4, 5,  1, 2]
// diff   = [-2, -2, -2, 3, 3] -> if diff < 0 -> can't be as start position
// from index 0 start to search

// if total amount < 0 -> return - 1
// diff -> [-2, -2, -2, 3, 3]
//.          ^ < 0             -> next
// diff -> [-2, -2, -2, 3, 3]
//.              ^ <           -> next
// diff -> [-2, -2, -2, 3, 3]
//.                 ^ <        -> next
// diff -> [-2, -2, -2, 3, 3]
//.                     ^ <    may be a start -> use amount to record -> amount += diff[i]
// diff -> [-2, -2, -2, 3, 3]
//.                        ^ < check amount += diff[i] -> if amount >= 0 until end -> this position is start
// ...
