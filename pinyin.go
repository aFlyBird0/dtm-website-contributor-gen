package main

import (
	"regexp"
	"strings"

	"github.com/mozillazg/go-pinyin"
)

func strAllLetter(str string) bool {
	match, _ := regexp.MatchString(`^[A-Za-z]+$`, str)
	return match
}

// 把中文名字转换为拼音
// 支持三个字以内的名字，否则请直接手动修改 csv 中的 「证书英文名」字段
func nameToPinyin(name string) string {
	if strings.TrimSpace(name) == "" {
		return ""
	}
	// 无中文字符就直接返回
	if strAllLetter(name) {
		return name
	}
	args := pinyin.NewArgs()
	pinyinSlice := pinyin.Pinyin(name, args)
	switch len(pinyinSlice) {
	case 1:
		return pinyinSlice[0][0]
	case 2:
		first := pinyinSlice[1][0]
		last := pinyinSlice[0][0]
		return firstUpper(first) + " " + firstUpper(last)
	case 3:
		first := pinyinSlice[1][0] + pinyinSlice[2][0]
		last := pinyinSlice[0][0]
		return firstUpper(first) + " " + firstUpper(last)
	default:
		return name
	}
}

// 首字母大写
func firstUpper(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
