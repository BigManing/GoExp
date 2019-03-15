package algorithm

import (
	"fmt"
	"testing"
)

// 普通 快速排序
func TestQickSortNormal(t *testing.T) {
	line()

	ints := []int{43, 6, 22, 447, 2, 888, 23, 466, 97}

	fmt.Println(ints)
	quickSort_normal(ints, 0, len(ints)-1)
	fmt.Println(ints)
}

// 普通 快速排序
func TestQickSortGoroution(t *testing.T) {

	ints := []int{43, 6, 22, 447, 2, 888, 23, 466, 97}

	fmt.Println(ints)
	ints = quickSort_goroutine(ints)

	fmt.Println(ints)
}

// 打印 线
func line() {
	println("===========================" +
		"" +
		"==============================")
}
