package main

import (
	"fmt"
	"sync"
)

// 问题描述
// 使用两个 goroutine 交替打印序列，一个 goroutine 打印数字， 另外一个 goroutine 打印字母， 最终效果如下：
// 12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728

func main() {
	letter, number := make(chan struct{}), make(chan struct{})
	wg := sync.WaitGroup{}

	go func() {
		i := 1
		for {
			select {
			case <-number:
				fmt.Println(i)
				i++
				fmt.Println(i)
				i++
				letter <- struct{}{}
			}
		}
	}()

	wg.Add(1)

	go func() {
		i := 'A'
		for {
			select {
			case <-letter:
				if i >= 'Z' {
					wg.Done()
					return
				}
				fmt.Println(string(i))
				i++
				fmt.Println(string(i))
				i++
				number <- struct{}{}
			}
		}
	}()
	number <- struct{}{}
	wg.Wait()
}
