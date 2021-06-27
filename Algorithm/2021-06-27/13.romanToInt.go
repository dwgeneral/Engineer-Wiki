package main

/*
 * 罗马数字由 I,V,X,L,C,D,M 构成；
   当小值在大值的左边，则减小值，如 IV=5-1=4；
   当小值在大值的右边，则加小值，如 VI=5+1=6；
*/
var symbolValues = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

func romanToInt(s string) (ans int) {
	n := len(s)
	for i := range s {
		value := symbolValues[s[i]]
		if i < n-1 && value < symbolValues[s[i+1]] {
			ans -= value
		} else {
			ans += value
		}
	}
	return
}
