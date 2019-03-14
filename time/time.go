package time

import (
	"fmt"
	"sync"
	"time"
)

var (
	working chan int //goroutine计数器 用于限制最大并发数
	wg      sync.WaitGroup
)

//https://www.jianshu.com/c/b03a9331735f
// 10个任务  在某个时间段内  只能处理三个
func handlWork(lists []int) {
	// 10个任务

	//println(lists[2:len(lists)])
	//	 定义chan  阻塞  三个
	//working = make(chan int, 3)
	//// 向通道 添加信息   超过3个会发生阻塞  这个地方可以控制后台同时处理 任务的个数
	//working <- k
	ticker := time.NewTicker(time.Second)
	//处理任务

	//每秒 提示一下
	maxCount := 3
	currentIndx := 0
	for {
		if currentIndx >= len(lists) {
			ticker.Stop()
			break
		}
		select {
		case <-ticker.C:

			// 角标
			// 任务分段
			for _, v := range lists[currentIndx : currentIndx+maxCount] {
				//wg.Add(1)
				go work(v)
			}
			currentIndx += maxCount
			break
		}

	}
	//wg.Wait()

}
func work(j int) {
	//defer wg.Done()

	fmt.Println("doing work#", j)
	//<-time.After(5 * time.Second) //假设5秒完成
	//
	////工作完成后计数器减1
	//<-working
}
