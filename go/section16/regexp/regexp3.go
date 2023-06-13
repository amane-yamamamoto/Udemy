package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`(abc|ABC)(xyz|XYZ)`)

	// 正規表現にマッチした文字列の取得
	// FindString
	fs1 := re.FindString("AAABCXYZZZ")
	fmt.Println(fs1)
	fs2 := re.FindAllString("ABCXYZABCXYZ", -1)
	fmt.Println(fs2)

	// 正規表現のグループによるサブマッチ
	// FindAllStringSubmatch
	re1 := regexp.MustCompile(`(\d+)-(\d+)-(\d+)`)
	s := `
			0123-456-7889
			111-222-333
			556-787-899
			`
	ms := re1.FindAllStringSubmatch(s, -1)
	fmt.Println(ms)

	for _, v := range ms {
		fmt.Println(v)
	}

	fmt.Println(ms[0][0])
	fmt.Println(ms[0][1])
	fmt.Println(ms[0][2])
}