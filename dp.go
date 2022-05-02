package DP

import "log"

func init() {
	log.SetFlags(0)
	log.SetPrefix("")
}

// 1863 Sum of All Subset XOR Totals
func subsetXORSum(nums []int) int {

	var PowerSet func(int) [][]int
	PowerSet = func(i int) [][]int {
		if i == len(nums) {
			return [][]int{{}}
		}
		r := [][]int{}
		for _, l := range PowerSet(i + 1) {
			r = append(r, l)
			r = append(r, append([]int{nums[i]}, l...))
		}
		return r
	}

	tSum := 0
	for _, l := range PowerSet(0) {
		if len(l) > 0 {
			v := l[0]
			for _, n := range l[1:] {
				v ^= n
			}
			tSum += v
		}
	}
	return tSum
}
