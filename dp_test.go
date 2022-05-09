package DP

import "log"

func init() {
	log.Print("> Dynamic Programming")
}

// 1863 Sum of All Subset XOR Totals
func Test1863(t *testing.T) {
	// 2^n Power sets
	var PowerSet func(i, N int) [][]int
	PowerSet = func(i, N int) [][]int {
		if i == N {
			return [][]int{{}}
		}

		r := [][]int{}
		for _, l := range PowerSet(i+1, N) {
			r = append(r, l)
			r = append(r, append([]int{i}, l...))
		}
		return r
	}

	// iterative PowerSet generator
	iPowerSet := func(N int) [][]int {
		r := [][]int{}
		P := 1
		for range N {
			P *= 2
		}
		for n := range P {
			l := []int{}
			for i := 0; n > 0; i++ {
				if n&1 == 1 {
					l = append(l, i)
				}
				n >>= 1
			}
			r = append(r, l)
		}
		return r
	}

	log.Printf("%v ---PowerSet-> %v", 4, iPowerSet(4))

	log.Printf("%v ---PowerSet-> %v", []int{0, 1, 2, 3}, PowerSet(0, 4))
	AAs := []string{"A", "G", "T", "C"}
	r := make([][]string, 16)
	for i, l := range PowerSet(0, 4) {
		for _, n := range l {
			r[i] = append(r[i], AAs[n])
		}
	}
	log.Printf("[A G T C] ---PowerSet-> %v", r)

	log.Print("6 ?= ", subsetXORSum([]int{1, 3}))
	log.Print("28 ?= ", subsetXORSum([]int{5, 1, 6}))
	log.Print("480 ?= ", subsetXORSum([]int{3, 4, 5, 6, 7, 8}))
}

// 79m Subsets
func Test79(t *testing.T) {
	Iterative := func(nums []int) [][]int {
		r := [][]int{{}}
		for _, n := range nums {
			t := append([][]int{}, r...)
			for _, l := range r {
				t = append(t, append(l, n))
			}
			r = t
		}
		return r
	}

	BackTrack := func(nums []int) [][]int {
		r := [][]int{}

		v := []int{}
		var kSet func(start, k int)
		kSet = func(start, k int) {
			if len(v) == k {
				r = append(r, append([]int{}, v...))
				return
			}
			for i := start; i < len(nums); i++ {
				v = append(v, nums[i])
				kSet(i+1, k)
				v = v[:len(v)-1]
			}
		}

		for k := 0; k <= len(nums); k++ {
			kSet(0, k)
		}
		return r
	}

	for _, f := range []func([]int) [][]int{subsets, Iterative, BackTrack} {
		log.Print("📀 ", runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name())
		log.Print("PowerSet-> ", f([]int{1, 2, 3}))
		log.Print("PowerSet-> ", f([]int{0}))
	}
}

// 131m Palindrome Partitioning
func Test131(t *testing.T) {
	log.Print(" -> ", partition("AAB"))
	log.Print(" -> ", partition("AABBAC"))
}

// 2597m The Number of Beautiful Subsets
func Test2597(t *testing.T) {
	log.Print("4 ?= ", beautifulSubsets([]int{2, 4, 6}, 2))
}
