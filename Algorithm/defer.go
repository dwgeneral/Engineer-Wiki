package main

import "fmt"

/*
defer 关键字的实现跟go关键字很类似，不同的是它调用的是runtime.deferproc而不是runtime.newproc。

在defer出现的地方，插入了指令call runtime.deferproc，然后在函数返回之前的地方，插入指令call runtime.deferreturn。

goroutine的控制结构中，有一张表记录defer，调用runtime.deferproc时会将需要defer的表达式记录在表中，而在调用runtime.deferreturn的时候，则会依次从defer表中出栈并执行。

因此，题目最后输出顺序应该是defer 定义顺序的倒序。panic 错误并不能终止 defer 的执行。
*/
func main() {
	defer_call()
}

func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
}

/*
defer 在定义的时候会计算好调用函数的参数，所以会优先输出10、20 两个参数。然后根据定义的顺序倒序执行。
*/
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}

/*
10, 1, 2, 3
20, 0, 2, 2
2, 0, 2, 2
1,1,3,4
*/
