package DP

import (
	"log"
	"math"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("")
}

// 22m Generate Parentheses
func generateParenthesis(n int) []string {
	r := []string{}

	var BT func(s string, opn, cls int)
	BT = func(s string, opn, cls int) {
		if opn == 0 && cls == 0 {
			r = append(r, s)
			return
		}

		if opn > 0 {
			BT(s+"(", opn-1, cls)
		}
		if cls > opn {
			BT(s+")", opn, cls-1)
		}
	}
	BT("", n, n)

	return r
}

// 39m Combination Sum
func combinationSum(candidates []int, target int) [][]int {
	r := [][]int{}

	var BT func(v []int, start, curSum int)
	BT = func(v []int, start, curSum int) {
		if curSum >= target {
			if curSum == target {
				r = append(r, append([]int{}, v...))
			}
			return
		}

		for i := start; i < len(candidates); i++ {
			BT(append([]int{candidates[i]}, v...), i, curSum+candidates[i])
		}
	}

	BT([]int{}, 0, 0)
	return r
}

// 46m Permutations
func permute(nums []int) [][]int {
	p := [][]int{}

	var BT func(v, s []int, l int)
	BT = func(v, s []int, l int) {
		if l == len(nums) {
			p = append(p, append([]int{}, v...))
			return
		}

		for i := 0; i < len(s); i++ {
			v = append(v, s[i])

			x := make([]int, len(s)-1)
			copy(x, s[:i])
			copy(x[i:], s[i+1:])

			BT(v, x, l+1)

			v = v[:len(v)-1]
		}
	}

	BT([]int{}, nums, 0)

	return p
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

// 279m Perfect Squares
func numSquares(n int) int {
	Mem := make([]int, n+1)
	for r := 1; r*r <= n; r++ {
		Mem[r*r] = 1
	}

	var Calc func(int) int
	Calc = func(n int) int {
		if Mem[n] > 0 {
			return Mem[n]
		}

		v := 4
		for r := 1; r*r < n; r++ {
			v = min(v, Calc(n-r*r)+Calc(r*r))
		}

		Mem[n] = v
		return v
	}

	return Calc(n)
}

// 552h Student Attendance Record II
func checkRecord(n int) int {
	D := make([][2][3]int, n+1) // {index, total-absence, consecutive-lateness}
	m := 1000_000_007

	D[0][0][0] = 1

	for i := 0; i < n; i++ {
		for a := 0; a < 2; a++ {
			for l := 0; l < 3; l++ {
				// 'P'
				D[i+1][a][0] += D[i][a][l]
				D[i+1][a][0] %= m

				// 'A'
				if a < 1 {
					D[i+1][a+1][0] += D[i][a][l]
					D[i+1][a+1][0] %= m
				}

				// 'L'
				if l < 2 {
					D[i+1][a][l+1] += D[i][a][l]
					D[i+1][a][l+1] %= m
				}
			}
		}
	}

	v := 0
	for a := 0; a < 2; a++ {
		for l := 0; l < 3; l++ {
			v += D[n][a][l]
			v %= m
		}
	}
	return v
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
