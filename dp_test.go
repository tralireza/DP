package DP

import (
	"log"
	"reflect"
	"runtime"
	"testing"
)

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

// 1255h Maximum Score Words Formed by Letters
func Test1255(t *testing.T) {
	WithPrune := func(words []string, letters []byte, score []int) int {
		x := 0

		v, vscore, vfrq := []string{}, 0, [26]int{}
		for _, l := range letters {
			vfrq[l-'a']++
		}

		var Walk func(int)
		Walk = func(start int) {
			if start == len(words) {
				log.Printf("%2d   %v   %q", vscore, vfrq, v)
				x = max(x, vscore)
				return
			}

			lfrq := [26]int{}
			copy(lfrq[:], vfrq[:])
			wscore := 0
			for _, r := range words[start] {
				if lfrq[r-'a'] == 0 {
					wscore = 0
					break
				}
				lfrq[r-'a']--
				wscore += score[r-'a']
			}
			if wscore == 0 {
				log.Printf("-> [P] %s %q", words[start], v)
				return
			}

			for i := start; i < len(words); i++ {
				v = append(v, words[start])
				vscore += wscore
				lfrq, vfrq = vfrq, lfrq

				Walk(i + 1)

				// BackTracking
				v = v[:len(v)-1]
				vscore -= wscore
				lfrq, vfrq = vfrq, lfrq
			}
		}

		for i := 0; i < len(words); i++ {
			Walk(i)
		}

		return x
	}

	for _, f := range []func([]string, []byte, []int) int{maxScoreWords, WithPrune} {
		log.Print("23 ?= ", f([]string{"dog", "cat", "dad", "good"}, []byte{'a', 'a', 'c', 'd', 'd', 'd', 'g', 'o', 'o'}, []int{1, 0, 9, 5, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}))
		log.Print("27 ?= ", f([]string{"xxxz", "ax", "bx", "cx"}, []byte{'z', 'a', 'b', 'c', 'x', 'x', 'x'}, []int{4, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 10}))
	}
}

// 2597m The Number of Beautiful Subsets
func Test2597(t *testing.T) {
	WithBitset := func(nums []int, k int) int {
		var Walk func(start, mask int) int
		Walk = func(start, mask int) int {
			if start == len(nums) {
				log.Print(" -> ", mask)
				if mask != 0 {
					return 1
				}
				return 0
			}

			g := true // Beautiful
			for i := 0; i < start && g; i++ {
				if (1<<i)&mask != 0 && (nums[i]-nums[start] == k || nums[start]-nums[i] == k) {
					g = false
				}
			}

			gn := 0
			if g {
				gn = Walk(start+1, mask+(1<<start))
			}
			return gn + Walk(start+1, mask)
		}

		return Walk(0, 0)
	}

	for _, f := range []func([]int, int) int{beautifulSubsets, WithBitset} {
		log.Print("49 ?= ", f([]int{2, 8, 4, 1, 7, 9, 5}, 3))
		log.Print("4 ?= ", f([]int{2, 4, 6}, 2))
	}
}
