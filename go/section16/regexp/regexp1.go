package main

import (
	"fmt"
	"regexp"
)

func main() {
	// regexp.MatchString
	match, _ := regexp.MatchString("A", "ABC")
	fmt.Println(match)


	// コンパイル
	re1, _ := regexp.Compile("A")
	match = re1.MatchString("ABC")
	fmt.Println(match)

	// MustCompile
	// エラーを返さず、直接 runtime errorを発生させる
	// 静的なものであればこちらを使うべき
	re2 := regexp.MustCompile("A")
	match = re2.MatchString("ABC")
	fmt.Println(match)

	// エスケープ
	// regexp.MustCompile("\\d")
	// regexp.MustCompile(`\d`)
}