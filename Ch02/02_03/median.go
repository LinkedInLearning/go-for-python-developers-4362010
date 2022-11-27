// Calculate the median of a slice of numbers
package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []float64{2, 1, 3}

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
