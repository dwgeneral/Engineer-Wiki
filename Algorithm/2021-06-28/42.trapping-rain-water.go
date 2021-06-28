package algorithm

/*
 * 方法一: 暴力法 时间 O(n^2) 空间 O(1)
 * 通过观察我们可以发现, 针对于坐标为 i 的点, 它的接水高度为 res = min(max(h[0]..h[i-1]), max(h[i+1..n])) - h[i], 题目求所有 res 之和
 */
func trap(height []int) (res int) {
	n := len(height)
	if n == 0 {
		return 0
	}

	for i := 1; i < n-1; i++ {
		// get leftMax
		leftMax := 0
		for j := 0; j <= i; j++ {
			if height[j] > leftMax {
				leftMax = height[j]
			}
		}
		// get RightMax
		rightMax := 0
		for j := i; j < n; j++ {
			if height[j] > rightMax {
				rightMax = height[j]
			}
		}
		res += min(leftMax, rightMax) - height[i]
	}
	return
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

/*
 * 方法二: 动态规划优化 时间 O(n) 空间 O(n)
 * 刚才的暴力法存在不必要的重复计算问题. 即求 i 点左右两边最高的点的高度, 我们只需要每次把左右两边的最高高度记录
   在两个数组中, 然后只需要和 h[i] 进行比较就可以了, 省去了很多重复计算
*/
func trap(height []int) (res int) {
	n := len(height)
	if n == 0 {
		return 0
	}
	leftMax := make([]int, n)
	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}
	rightMax := make([]int, n)
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}
	for i := 0; i < n; i++ {
		res += min(leftMax[i], rightMax[i]) - height[i]
	}
	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

/*
 * 方法三: 双指针 时间 O(n) 空间 O(1)
 * 我们注意到方法二中的数组只与上一个值相关, 一般这种动归都可以进行状态压缩, 而且从左右两边用两个数组分别记录较大值, 所以想到可以使用左右两个指针来做, 并用两个变量分别记录左右较大值
 * 具体的, 两个指针不断向中间靠近, 如果 h[l] < h[r], 则 l 点的接水量就可以确定了, 那么 l++, 反之 r--
 */
func trap(height []int) (res int) {
	n := len(height)
	if n == 0 {
		return 0
	}
	left, right, leftMax, rightMax := 0, n-1, 0, 0
	for left < right {
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])
		if height[left] < height[right] {
			res += leftMax - height[left]
			left++
		} else {
			res += rightMax - height[right]
			right--
		}
	}
	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
