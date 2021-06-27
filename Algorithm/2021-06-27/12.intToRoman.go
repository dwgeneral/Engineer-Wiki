package main

/*
这个思路相对比较简单，因为题目中说输入在 1 ~3999 的范围内，所以我们把 1 到 9，10 到 90，100 到 900，1000 到 3000 对应的罗马数字都表示出来，最后对于任何输入，我们要做的就是把找到的罗马数字组合起来即可。
比如输入是 2359，我们找到 2000，300，50，9 对应的罗马数字为 MM，CCC，L，IX，组合后得到结果为 MMCCCLIX
*/
func intToRoman(num int) string {
	var (
		thousands = []string{"", "M", "MM", "MMM"}
		hundreds  = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
		tens      = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
		ones      = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	)
	return thousands[num/1000] + hundreds[num%1000/100] + tens[num%100/10] + ones[num%10]
}
