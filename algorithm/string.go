package algorithm

import (
	"strings"
)

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

//输出字符串中第一个只出现一次的字符，用两种方案。
func getSingleCharFromStr1(str string) string {
	//	思路1  遍历str  存入map   保存出现的次数
	//再次遍历  查找第一个出现一次的 char
	keyMap := make(map[string]int)
	for _, v := range str {
		keyMap[string(v)] += 1
	}

	for _, v := range str {
		if keyMap[string(v)] == 1 {
			return string(v)
		}
	}
	return ""
}

//输出字符串中第一个只出现一次的字符，用两种方案。
func getSingleCharFromStr2(str string) string {
	//	思路2  遍历str
	for _, v := range str {
		//if  i == strings.LastIndexAny(str, string(v)) {  // 这个操作是错误的    会返回 g

		//  如果  字符第一次出现的索引 == 最后一次出现的索引
		if strings.Index(str, string(v)) == strings.LastIndexAny(str, string(v)) {
			return string(v)
		}
	}
	return ""
}
