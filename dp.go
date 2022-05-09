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

// 78m Subsets
func subsets(nums []int) [][]int {
	var PowerSet func([]int) [][]int
	PowerSet = func(v []int) [][]int {
		if len(v) == 0 {
			return [][]int{{}}
		}

		r := [][]int{}
		for _, l := range PowerSet(v[1:]) {
			r = append(r, l)
			r = append(r, append([]int{v[0]}, l...))
		}
		return r
	}

	return PowerSet(nums)
}

// 131m Palindrome Partitioning
func partition(s string) [][]string {
	r := [][]string{}

	Palindrome := func(s string) bool {
		l, r := 0, len(s)-1
		for l < r {
			if s[l] != s[r] {
				return false
			}
			l++
			r--
		}
		return true
	}

	var v []string
	var Walk func(int)
	Walk = func(i int) {
		if i == len(s) {
			r = append(r, append([]string{}, v...))
			return
		}

		for j := i + 1; j <= len(s); j++ {
			if Palindrome(s[i:j]) {
				v = append(v, s[i:j])
				Walk(j)
				v = v[:len(v)-1]
			}
		}
	}

	Walk(0)
	return r
}

// 2597m The Number of Beautiful Subsets
func beautifulSubsets(nums []int, k int) int {
	r := [][]int{}

	v := []int{}
	var Walk func(int)
	Walk = func(start int) {
		if start == len(nums) {
			for l := 0; l < len(v); l++ {
				for r := l + 1; r < len(v); r++ {
					if v[l]-v[r] == k || v[r]-v[l] == k {
						return
					}
				}
			}
			r = append(r, append([]int{}, v...))
			return
		}

		for i := start; i < len(nums); i++ {
			v = append(v, nums[start])
			Walk(i + 1)
			v = v[:len(v)-1]
		}
	}
	for start := range nums {
		Walk(start)
	}

	log.Print(r)
	return len(r)
}
