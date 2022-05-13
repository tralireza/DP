package DP

import (
	"log"
	"math"
)

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

// 1255h Maximum Score Words Formed by Letters
func maxScoreWords(words []string, letters []byte, score []int) int {
	xScore := 0

	v := []string{}
	var Walk func(int)
	Walk = func(start int) {
		if start == len(words) {
			lFQ := [26]int{}
			for _, l := range letters {
				lFQ[l-'a']++
			}

			vScore := 0
			for _, w := range v {
				for i := 0; i < len(w); i++ {
					if lFQ[w[i]-'a'] == 0 {
						return
					}
					lFQ[w[i]-'a']--
					vScore += score[w[i]-'a']
				}
			}
			xScore = max(xScore, vScore)
			return
		}

		for i := start; i < len(words); i++ {
			v = append(v, words[start])
			Walk(i + 1)
			v = v[:len(v)-1]
		}
	}

	for i := 0; i < len(words); i++ {
		Walk(i)
	}

	return xScore
}

// 2597m The Number of Beautiful Subsets
func beautifulSubsets(nums []int, k int) int {
	r := [][]int{}

	v := []int{}
	var Walk func(int)
	Walk = func(start int) {
		if start == len(nums) {
			log.Print(" -> ", v)
			r = append(r, append([]int{}, v...))
			return
		}

		g := true
		for i := 0; i < len(v) && g; i++ {
			g = v[i]-nums[start] != k && nums[start]-v[i] != k
		}
		if !g {
			log.Print(v, nums[start], " -> P")
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

// 3068h Find the Maximum Sum of Node Values
func maximumValueSum(nums []int, k int, edges [][]int) int {
	M := make([][2]int, len(nums)+1)

	M[len(nums)][0] = 0           // even XOR operations
	M[len(nums)][1] = math.MinInt // odd XOR operations

	for i := len(nums) - 1; i >= 0; i-- {
		for parity := 0; parity < 2; parity++ {
			xor := nums[i] ^ k + M[i+1][parity^1]
			noXor := nums[i] + M[i+1][parity]

			M[i][parity] = max(xor, noXor)
		}
	}

	log.Print(M)
	return M[0][0]
}
