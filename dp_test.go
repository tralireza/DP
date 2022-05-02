package DP

import "log"

func init() {
	log.Print("> Dynamic Programming")
}

// 1863 Sum of All Subset XOR Totals
func Test1863(t *testing.T) {
	// 2^n Power sets
	var PowerSet func([]int, int) [][]int
	PowerSet = func(nums []int, i int) [][]int {
		if i == len(nums) {
			return [][]int{{}}
		}

		r := [][]int{}
		for _, l := range PowerSet(nums, i+1) {
			r = append(r, l)
			r = append(r, append([]int{nums[i]}, l...))
		}
		return r
	}
	log.Printf("%v -> %v", []int{0, 1, 2, 3}, PowerSet([]int{0, 1, 2, 3}, 0))

	log.Print("6 ?= ", subsetXORSum([]int{1, 3}))
	log.Print("28 ?= ", subsetXORSum([]int{5, 1, 6}))
	log.Print("480 ?= ", subsetXORSum([]int{3, 4, 5, 6, 7, 8}))
}
