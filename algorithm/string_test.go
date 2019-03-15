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
