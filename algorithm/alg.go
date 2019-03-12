package main

import (
	"fmt"
	"sync"
)

func main() {
	ints := []int{43, 6, 22, 447, 2, 888, 23, 466, 97}

	fmt.Println(ints)
	quickSort_normal(ints, 0, len(ints)-1)
	//ints = quickSort_goroutine(ints)

	fmt.Println(ints)
}

//  一般 快速排序
// 在原有的数组的基础上  进行排序
func quickSort_normal(datas []int, begin int, end int) {
	if begin < end {
		flag := datas[(begin+end)/2]
		p1, p2 := begin, end
		for p1 <= p2 {
			// 找到 左边 大的 index
			for datas[p1] < flag {
				p1++
			}
			//找到右边 小的 index
			for datas[p2] > flag {
				p2--
			}
			//位置互换
			if p1 <= p2 {
				datas[p1], datas[p2] = datas[p2], datas[p1]
				p1++
				p2--
			}
		}
		// 递归
		if begin < p2 {
			quickSort_normal(datas, begin, p2)
		}
		if end > p1 {
			quickSort_normal(datas, p1, end)
		}
	}
}

//	思路  使用两个slice   小的  大的    最后 在组合起来
func quickSort_goroutine(datas []int) []int {
	// 条件判断
	if len(datas) < 1 {
		return datas
	}
	flag := datas[0]
	// 基于 flag  就行分组
	var s1, s2 []int
	for k, v := range datas {
		if k == 0 {
			continue
		}
		if v > flag {
			s2 = append(s2, v)
		}
		if v < flag {
			s1 = append(s1, v)
		}
	}

	// 开启 goroutine   各自再排序
	sw := sync.WaitGroup{}
	sw.Add(2)
	go func() {
		s2 = quickSort_goroutine(s2)
		sw.Done()
	}()
	go func() {
		s1 = quickSort_goroutine(s1)
		sw.Done()
	}()
	sw.Wait()
	//
	return append(append(s1, flag), s2...)
}
