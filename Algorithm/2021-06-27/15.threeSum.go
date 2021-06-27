package main

import "sort"

func threeSum(nums []int) (res [][]int) {
	n := len(nums)
	if n < 3 {
		return
	}
	sort.Ints(nums)
	for index, value := range nums {
		if nums[index] > 0 {
			return
		}
		if index > 0 && nums[index] == nums[index-1] {
			continue
		}
		l, r := index+1, n-1
		for l < r {
			sum := value + nums[l] + nums[r]
			switch {
			case sum == 0:
				res = append(res, []int{value, nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			case sum > 0:
				r--
			case sum < 0:
				l++
			}
		}
	}
	return
}
