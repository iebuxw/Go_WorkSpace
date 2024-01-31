package main

import (
	"fmt"
	"strings"
	"crypto/md5"
	"encoding/hex"
)

// 格式化字符串
func Sprintf(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a...)
}

// 计算字符串的长度
func Strlen(s string) int {
	return len(s)
}

// 替换字符串中的字符 str_replace
func StrReplace(s, old, new string) string {
	return strings.Replace(s, old, new, -1) // -1代表全部替换
}

// 返回字符串的一部分
func Substr(s string, start, length int) string {
	return s[start:start+length]
}

// 返回中文字符串的一部分 mb_substr
func MbSubstr(s string, start, length int) string {
	rs := []rune(s)
	return string(rs[start:start+length])
}

// 首次出现位置，没找到返回-1
func Strpos(s, substr string) int {
	return strings.Index(s, substr)
}

// 首次出现位置并返回str1结尾的字符串，未出现返回false
func Strstr(s, substr string) string {
	index := strings.Index(s, substr)
	if index == -1 {
		return ""
	}
	return s[index:]
}

// 将字符串转换为小写
func StrToLower(s string) string {
	return strings.ToLower(s)
}

// 将字符串转换为大写
func StrToUpper(s string) string {
	return strings.ToUpper(s)
}

// 去除字符串两端的空格
func Trim(s string) string {
	return strings.TrimSpace(s)
}

// 把字符串打散为切片
func Explode(sep, s string) []string {
	return strings.Split(s, sep)
}

// 计算字符串的 MD5 散列值
func Md5(s string) string {
	hasher := md5.New()
	hasher.Write([]byte(s))
	return hex.EncodeToString(hasher.Sum(nil))
}

func main() {
	fmt.Println("Sprintf == ", Sprintf("Hello, %s!", "world"))          // 格式化字符串
	fmt.Println("Strlen == ", Strlen("Hello, world"))                  // 计算字符串的长度
	fmt.Println("StrReplace == ", StrReplace("Hello, world", "world", "Go")) // 替换字符串中的字符
	fmt.Println("Substr == ", Substr("Hello, world", 7, 5))            // 返回字符串的一部分
	fmt.Println("MbSubstr == ", MbSubstr("你好，世界", 0, 2))              // 返回中文字符串的一部分
	fmt.Println("Strpos == ", Strpos("Hello, world", "ee"))              // 首次出现位置
	fmt.Println("Strstr == ", Strstr("Hello, world", "o"))              // 首次出现位置并返回str1结尾的字符串
	fmt.Println("StrToLower == ", StrToLower("Hello, WORLD"))               // 将字符串转换为小写
	fmt.Println("StrToUpper == ", StrToUpper("Hello, world"))               // 将字符串转换为大写
	fmt.Println("Trim == ", Trim("  Hello, world  "))                 // 去除字符串两端的空格
	fmt.Println("Explode == ", Explode(",", "apple,banana,orange"))      // 把字符串打散为数组
	fmt.Println("Md5 == ", Md5("Hello, world"))                      // 计算字符串的 MD5 散列值
}
