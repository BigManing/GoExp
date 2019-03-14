package algorithm

import "strings"

// 字符串中   最长的无重复  子串

// 取出 所有可能的字符串
func getLongestString(strs string) (maxLen int, str string) {
	length := len(strs)
	for x := 0; x < length; x++ {
		for y := x + 1; y <= length; y++ {
			//	 查看  x-y 中间的string  是否有重复的
			tmpStr := strs[x:y]
			if checkRangeStr(tmpStr) && len(tmpStr) > maxLen {
				maxLen = len(tmpStr)
				str = tmpStr
			}
		}
		println("-----------------------------")
	}
	return
}

//检查是否满足条件  （不存在重复字符）
func checkRangeStr(tmpStr string) bool {
	println("原值=", tmpStr)
	for k, v := range tmpStr {
		s := tmpStr[0:k]
		//println("取值=",s)
		if strings.Contains(s, string(v)) {
			return false
		}
	}
	return true
}
