package _go

/*
sync.WaitGroup内部维护着一个计数器，计数器的值可以增加和减少。
例如当我们启动了N 个并发任务时，就将计数器值增加N。每个任务完成时通过调用Done()方法将计数器减1。
通过调用Wait()来等待并发任务执行完，当计数器值为0时，表示所有并发任务已经完成。
*/

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello() {
	defer wg.Done()
	fmt.Println("Hello Goroutine!")
}
func main() {
	wg.Add(1)
	go hello() // 启动另外一个goroutine去执行hello函数
	fmt.Println("main goroutine done!")
	wg.Wait()
}
