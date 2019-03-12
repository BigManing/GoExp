package main

import (
	"fmt"
)

// slice  是一个拥有长度
func main() {
	//    getLenCap()

}

// 这个操作 对应的打印结果
func getLenCap() {
	s := []int{5}
	// fmt.Println(cap(s)) //1

	s = append(s, 7)
	// fmt.Println(cap(s)) //2

	s = append(s, 9)
	// fmt.Println(cap(s)) //4

	x := append(s, 11)
	//  此时 还是共用的同一个数组的 [5 7 9 11]

	y := append(s, 12) // 在s的基础上 append 12    array[3]=12    原先11 被修改了

	//  看到结果意不意外  惊不惊喜
	fmt.Println(s, x, y) //[5 7 9] [5 7 9 12] [5 7 9 12]
}
