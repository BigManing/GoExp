package algorithm

import (
	"fmt"
	"testing"
)

//  取出字符串中  最长的不重复的 子串
func TestString(t *testing.T) {
	line()

	strs := "abcdddcccc"
	maxLen, str := getLongestString(strs)
	fmt.Println(maxLen, str)
}

func TestGetSingleCharFromStr1(t *testing.T) {
	str := "good good study"
	result := "s"

	charFromStr1 := getSingleCharFromStr1(str)
	t.Log(charFromStr1)

	if charFromStr1 != result {
		t.Error("失败")
	} else {
		t.Log("成功")
	}
}
func TestGetSingleCharFromStr2(t *testing.T) {
	str := "good good study"
	result := "s"

	charFromStr1 := getSingleCharFromStr2(str)
	t.Log(charFromStr1)

	if charFromStr1 != result {
		t.Error("失败")
	} else {
		t.Log("成功")
	}

}
