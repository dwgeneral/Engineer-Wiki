package main

import "math"

/*
 * 一个数 /10 得到除个位外的数, %10 得到个位数, 先将负数统一转为正数处理
 */
func reverse(x int) int {
	sign, res := 1, 0
	if x < 0 {
		x = int(math.Abs(float64(x)))
		sign = -1
	}
	for x != 0 {
		if res > math.MaxInt32/10 {
			return 0
		}
		res = res*10 + (x % 10)
		x /= 10
	}

	return sign * res
}
