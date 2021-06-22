package main

import "fmt"

type student struct {
	Name string
	Age  int
}

func main() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
	}
	fmt.Printf("%v", m)
}

// golang 的 for ... range 语法中，stu 变量会被复用，每次循环会将集合中的值复制给这个变量，因此，会导致最后m中的map中储存的都是stus最后一个student的值
// map[li:0xc00000c060 wang:0xc00000c060 zhou:0xc00000c060]
