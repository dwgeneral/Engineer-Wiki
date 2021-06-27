package main

/*
 * 这题可以使用双指针, 类似于滑动窗口, 从左往右, 随着右指针不断扩大, 判断是否存在重复字符, 这里的判重可以使用 map 或者 数组来实现
 * 如果存在重复字符, 就不断缩小左边界, 依次记录窗口内字符的长度, 取最大值
 * 提前结束循环条件: 如果左指针 + 最大长度 已经大于 字符串长度了, 则可以直接返回结果了, 因为即便后边的字符都是不重复的, 也不会比当前最大长度大了
 * 时间 O(n) 空间 O(n)
 */
func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n < 2 {
		return n
	}
	bitMap := [256]bool{}
	left, right, res := 0, 0, 0
	for right < n {
		if bitMap[s[right]] {
			bitMap[s[left]] = false
			left++
		} else {
			bitMap[s[right]] = true
			right++
		}
		res = max(res, right-left)
		if left+res >= n {
			break
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
