package main

import (
	"fmt"
	"sync"
)

/*
一个标准的协程+信道实现
*/

func main() {



	taskChan := make(chan int)
	TCount := 10
	var wg sync.WaitGroup //创建一个sync.WaitGroup

	// 产生任务
	go func() {
		for i := 0; i < 10; i++ {
			taskChan <- i
		}
		// 全部任务都输入后关闭信道，告诉工作者进程没有新任务了。
		close(taskChan)
	}()

	// 告诉 WaitGroup 有 TCount 个执行者。
	wg.Add(TCount)
	// 启动 TCount 个协程执行任务
	for i := 0; i < TCount; i++ {

		// 注意：如果协程内使用了 i，必须有这一步，或者选择通过参数传递进协程。
		// 否则 i 会被 for 所在的协程修改，协程实际使用时值并不确定。
		i := i

		go func() {

			// 协程结束时报告当前协程执行完毕。
			defer func() { wg.Done() }()

			fmt.Printf("工作者 %v 启动...\r\n", i)

			for task := range taskChan {

				// 建立匿名函数执行任务的目的是为了捕获单个任务崩溃，防止造成整个工作者、系统崩溃。
				func() {

					defer func() {
						err := recover()
						if err != nil {
							fmt.Printf("任务失败：工作者i=%v, task=%v, err=%v\r\n", i, task, err)
						}
					}()

					// 故意崩溃，看看是不是会造成整个系统崩溃。
					if task%100==0{
						panic("故意崩溃啦")
					}

					// 这里的 task 并不需要通过参数传递进来。
					// 原因是这里是同步执行的，并不会被其它协程修改。
					fmt.Printf("任务结果=%v ，工作者id=%v, task=%v\r\n",task*task,i,task)
				}()
			}

			fmt.Printf("工作者 %v 结束。\r\n", i)
		}()
	}

	//等待所有任务完成
	wg.Wait()
	print("全部任务结束")
}