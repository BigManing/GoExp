package algorithm

import (
	"fmt"
	"sync"
)

//下面的迭代会有什么问题？

type threadSafeSet struct {
	sync.RWMutex
	s []interface{}
}

//  线程安全的map 一定要用锁
// 这个题目 考察的是 chan的缓冲
func (set *threadSafeSet) Iter() <-chan interface{} {
	//ch := make(chan interface{}) // 解除注释看看！
	ch := make(chan interface{}, len(set.s))
	go func() {
		set.RLock()

		for elem, value := range set.s {
			ch <- elem
			println("Iter:", elem, value)
		}

		close(ch)
		set.RUnlock()

	}()
	return ch
}

func printMap() {
	list = make(map[string]Test)
	name := Test{"xiaoming"}
	list["name"] = name
	//list["name"].Name = "Hello"  // 这段会报错    value 是不寻址的   如果改成string *Test就会不一样
	fmt.Println(list["name"])
}

var list map[string]Test

type Test struct {
	Name string
}
