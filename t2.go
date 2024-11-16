package main

import (
	"fmt"
)

func beautiful(nums []int) {
	var i, j, count int
	j = i + 1
	for p := 0; p < len(nums)-2; p++ {
		for i := 1; i < len(nums)-1; i++ {
			if nums[j] >= nums[i] {
				count = count + (nums[j] - nums[i])
			} else if nums[j] < nums[i] {
				count = count + (nums[i] - nums[j])
			}
		}
	}
	fmt.Println(count)
}

func main() {
	var a int
	fmt.Scan(&a)
	nums := make([]int, a)
	for i := 0; i < a; i++ {
		fmt.Scan(&nums[i])
	}
	beautiful(nums)
}
