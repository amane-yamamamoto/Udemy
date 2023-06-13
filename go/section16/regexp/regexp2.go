package main

import (
	"fmt"
	"regexp"
)

func main() {
	// 正規表現のフラグ
	/*
		i 大文字小文字を区別しない
		m マルチラインモード
		s .か\nにマッチ
		U 最小マッチへの変換（x*はx*?へ、x+はx+?へ）
	*/

	re3 := regexp.MustCompile(`(?i)abc`)
	match := re3.MatchString("ABC")
	fmt.Println(match)

	// 幅をもたない正規表現のパターン
	re4 := regexp.MustCompile(`^ABC$`)
	match = re4.MatchString("ABC")
	fmt.Println(match)

	match = re4.MatchString(" ABC  ")
	fmt.Println(match)

	// 繰り返しを表す正規表現
	re6 := regexp.MustCompile("a+b*")
	fmt.Println(re6.MatchString("ab"))
	fmt.Println(re6.MatchString("a"))
	fmt.Println(re6.MatchString("aaabbbbbb"))
	fmt.Println(re6.MatchString("b"))

	// 正規表現の文字クラス
	re8 := regexp.MustCompile(`[XYZ]`)
	fmt.Println(re8.MatchString("Y"))

	re9 := regexp.MustCompile(`^[0-9A-Za-z_]{3}$`)
	fmt.Println(re9.MatchString("ABC"))
	fmt.Println(re9.MatchString("abcdefg"))


	// 正規表現のグループ
	re11 := regexp.MustCompile(`(abc|ABC)(xyz|XYZ)`)
	fmt.Println(re11.MatchString("abcxyz"))
	fmt.Println(re11.MatchString("ABCXYZ"))
	fmt.Println(re11.MatchString("ABCxyz"))
	fmt.Println(re11.MatchString("ABCabc"))
}

