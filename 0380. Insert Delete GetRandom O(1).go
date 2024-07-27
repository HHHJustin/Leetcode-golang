package main

import "math/rand"

// We need a map and a array.
// array to save each element.
// map -> key = element -> value = the index.
type RandomizedSet struct {
	arr  []int
	size int
	set  map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		arr:  make([]int, 0),
		size: 0,
		set:  make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, OK := this.set[val]
	if OK {
		return false
	}
	this.set[val] = this.size
	this.size++
	this.arr = append(this.arr, val)
	return true
}

// if the val was present -> return false.
// if the item was not present -> map[val] = len(arr), arr.append -> return true.

func (this *RandomizedSet) Remove(val int) bool {
	index, OK := this.set[val]
	if !OK {
		return false
	}
	temp := this.arr[this.size-1]
	this.arr[this.size-1] = this.arr[index]
	this.arr[index] = temp
	this.set[temp] = index
	this.arr = this.arr[:this.size-1]
	this.size--
	delete(this.set, val)
	return true
}

// if the val was not present -> return false.
// if the val was present -> remove val in arr and map -> return true

func (this *RandomizedSet) GetRandom() int {
	n := rand.Intn(this.size)
	return this.arr[n]
}

// use rand.Intn() to get random index
// return the random element from the current set of element

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
