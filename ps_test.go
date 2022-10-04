package PrefixSum

import (
	"log"
	"testing"
)

func init() {
	log.Print("> Prefix Sum")
}

// 523m Continuous Subarray Sum
func Test523(t *testing.T) {
	log.Print("true ? ", checkSubarraySum([]int{23, 2, 4, 6, 7}, 6))
	log.Print("true ? ", checkSubarraySum([]int{23, 2, 6, 4, 7}, 6))
	log.Print("false ? ", checkSubarraySum([]int{23, 2, 6, 4, 7}, 13))
}

// 974m Subarray Sums Divisible by K
func Test974(t *testing.T) {
	log.Print("7 ?= ", subarraysDivByK([]int{4, 5, 0, -2, -3, 1}, 5))
	log.Print("0 ?= ", subarraysDivByK([]int{5}, 9))
}
