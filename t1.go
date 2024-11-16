package main

import (
	"fmt"
)

func max(nums []int) (a []int) {
	for i := 0; i < len(nums); i++ {
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] < nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
	}
	return a
}

func sleep(nums []int, n int) {
	var count int
	max(nums)
	for i, j := range nums {
		count += j
		if count > n {
			fmt.Println(i + 1)
			break
		}
	}
	if count < n {
		fmt.Println(-1)
	}
}

func main() {
	var n, q, a, b int
	fmt.Scan(&n, &q)
	num := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		num[i] = a
	}
	for i := 0; i < q; i++ {
		fmt.Scan(&b)
		sleep(num, b)
	}
}
