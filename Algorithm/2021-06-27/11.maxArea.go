package main

/*
 * 双指针, 一头一尾, 计算面积, 哪边高度小, 哪边移动, 终止条件为指针相遇
 * 时间 O(n)
 */
func maxArea(height []int) int {
	left, right, res := 0, len(height)-1, 0
	for left < right {
		area := min(height[left], height[right]) * (right - left)
		res = max(res, area)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return res
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
