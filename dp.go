package DP

import "log"

func init() {
	log.SetFlags(0)
	log.SetPrefix("")
}

// 1863 Sum of All Subset XOR Totals
func subsetXORSum(nums []int) int {
	rCalls := 0
	var Calc func(i, v int) int
	Calc = func(i, v int) int {
		rCalls++
		if i == len(nums) {
			return v
		}

		xSum := 0
		xSum += Calc(i+1, v)
		xSum += Calc(i+1, v^nums[i])
		return xSum
	}

	v := Calc(0, 0)
	log.Print("rCalls -> ", rCalls)
	return v
}
