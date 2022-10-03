package PrefixSum

import (
	"log"
)

func init() {
	log.SetFlags(0)
}

// 523m Continuous Subarray Sum
func checkSubarraySum(nums []int, k int) bool {
	/*
	   Prefix_i = a_0 + ... + a_i
	   Sum(Subarray[i+1, j]) = Prefix_j - Prefix_i
	   (Prefix_j - Prefix_i) % k == 0  =>  Prefix_j == Prefix_i (mod k)
	*/

	Index := map[int]int{0: -1}

	mpfx := 0
	for r := range nums {
		mpfx += nums[r]
		mpfx %= k

		if l, ok := Index[mpfx]; ok {
			if l+1 < r {
				return true
			}
		} else {
			Index[mpfx] = r
		}
	}

	return false
}
