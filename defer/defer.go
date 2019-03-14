package main

import "fmt"

func main() {
	f1()
}

func f2() {
	a := 1
	b := 2
	defer calc(a, b)
	a = 0
	defer calc(a, b)
	/**
	0 2 2
	1 2 3
	考察两个知识点：
	1.defer是栈调用，后写的先执行
	2.defer的函数调用语句会在父函数调用后执行，但是用到的【参数】会在当时就执行得出
	*/

}
func f1() {
	a := 1
	b := 2
	defer calc(a, calc(a, b)) //  defer 函数的参数 需要当时就计算  第二个参数位 calc(1, 2)=3 （ 内部打印1 2 3 ）  calc(1,3)然后入栈
	a = 0
	defer calc(a, calc(a, b)) // 同上 计算出 第二个 calc(0, 2)=2 （打印 0 2 2）  第一个calc(0, 2) 入栈

	//  函数 执行完成后  先执行 栈顶的函数

	/**
	1 2 3
	0 2 2
	0 2 2
	1 3 4

	考察两个知识点：
	1.defer是栈调用，后写的先执行
	2.defer的函数调用语句会在父函数调用后执行，但是用到的参数会在当时就执行得出
	*/

}

func calc(x, y int) int {
	fmt.Println(x, y, x+y)
	return x + y
}
