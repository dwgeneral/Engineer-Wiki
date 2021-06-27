package main

import "sort"

// 和三数之和基本一样的解法
func fourSum(nums []int, target int) (res [][]int) {
	n := len(nums)
	if n < 4 {
		return
	}
	sort.Ints(nums)

	for firstIdx := 0; firstIdx < n-3; firstIdx++ {
		if firstIdx > 0 && nums[firstIdx] == nums[firstIdx-1] {
			continue
		}
		for secondIdx := firstIdx + 1; secondIdx < n-2; secondIdx++ {
			if secondIdx > firstIdx+1 && nums[secondIdx] == nums[secondIdx-1] {
				continue
			}
			l, r := secondIdx+1, n-1
			for l < r {
				sum := nums[firstIdx] + nums[secondIdx] + nums[l] + nums[r]
				switch {
				case sum == target:
					res = append(res, []int{nums[firstIdx], nums[secondIdx], nums[l], nums[r]})
					for l < r && nums[l] == nums[l+1] {
						l++
					}
					for l < r && nums[r] == nums[r-1] {
						r--
					}
					l++
					r--
				case sum > target:
					r--
				case sum < target:
					l++
				}
			}
		}
	}
	return
}
