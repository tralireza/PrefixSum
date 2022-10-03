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
