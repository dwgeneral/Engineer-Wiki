//写代码实现两个 goroutine，其中一个产生随机数并写入到 go channel 中，另外一个从 channel 中读取数字并打印到标准输出。最终输出五个随机数。

package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	goch := make(chan int)

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			goch <- rand.Intn(10)
		}
		close(goch)

	}()

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := range goch {
			fmt.Println(i)
		}
	}()

	wg.Wait()

}
