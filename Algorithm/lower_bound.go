package main

import "fmt"

// 在一组有序数组中找到第一个大于等于target的元素, 返回其下标
func lowerBound(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}

func main() {
	nums := []int{2, 2}
	a := lowerBound(nums, 3)
	fmt.Println(a)
}
