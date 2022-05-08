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

	Mem := map[string]byte{}
	Palindrome := func(s string) bool {
		if v, ok := Mem[s]; ok {
			return v == 1
		}
		l, r := 0, len(s)-1
		for l < r {
			if s[l] != s[r] {
				Mem[s] = 0
				return false
			}
			l++
			r--
		}
		Mem[s] = 1
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

	for l := 1; l <= len(s); l++ {
		if Palindrome(s[0:l]) {
			v = []string{s[0:l]}
			Walk(l)
		}
	}

	log.Print(len(Mem), " -> ", Mem)
	return r
}
