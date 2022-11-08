// Calculate the median of a random list of numbers
package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	const size = 11
	var nums []float64
	for i := 0; i < size; i++ {
		nums = append(nums, rand.Float64()*100)
	}

	sort.Float64s(nums)
	var median float64
	i := len(nums) / 2
	if len(nums)%2 == 1 {
		median = nums[i]
	} else {
		median = (nums[i-1] + nums[i]) / 2
	}
	fmt.Println(median)
}
