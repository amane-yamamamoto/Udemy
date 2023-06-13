package main

import (
	"fmt"
	"strings"
)

func main() {

	// 文字列の結合
	s1 := strings.Join([]string{"A", "B", "C"}, ",")
	s2 := strings.Join([]string{"A", "B", "C"}, "")
	fmt.Println(s1, s2)

	// 部分文字列検索（インデックスを返す）
	i1 := strings.Index("ABCDE", "A")
	i2 := strings.Index("ABCDE", "D")
	fmt.Println(i1, i2)

	i3 := strings.LastIndex("ABCDABCD", "BC")
	fmt.Println(i3)

	i4 := strings.IndexAny("ABC", "ABC")
	fmt.Println(i4)

	// 指定した文字列で開始する　or 終わるか
	b1 := strings.HasPrefix("ABC", "A")
	b2 := strings.HasSuffix("ABC", "C")
	fmt.Println(b1, b2)

	// 含むか
	b3 := strings.Contains("ABC", "B") 
	b4 := strings.ContainsAny("ABCDE", "BD")
	fmt.Println(b3, b4)

	// どれだけ現れるか
	i6 := strings.Count("ABCABC", "B")
	i7 := strings.Count("ABCABC", "")
	fmt.Println(i6, i7)

	// 繰り返して結合する
	s3 := strings.Repeat("ABC", 4)
	s4 := strings.Repeat("ABC", 0)
	fmt.Println(s3, s4)

	// 置換
	s5 := strings.Replace("AAAAAA", "A", "B", 1)
	s6 := strings.Replace("AAAAAA", "A", "B", -1)
	fmt.Println(s5, s6)

	// 分離する
	s7 := strings.Split("A,B,C,D,E", ",")
	fmt.Println(s7)
	s8 := strings.SplitAfter("A,B,C,D,E", ",")
	fmt.Println(s8)
	s9 := strings.SplitN("A,B,C,D,E", ",", 2)
	fmt.Println(s9)
	s10 := strings.SplitAfterN("A,B,C,D,E", ",", 4)
	fmt.Println(s10)
	
	// 大文字小文字の変換
	s11 := strings.ToLower("ABC")
	s13 := strings.ToUpper("abc")
	fmt.Println(s11, s13)

	// 文字列から前後の空白を取り除く（間は取り除かれない）（全角も対象）
	s15 := strings.TrimSpace(" - ABC  - ")
	s16 := strings.TrimSpace("　　　ABC  ")
	fmt.Println(s15, s16)

	// 文字列からスペースで区切られたフィールドを取り出す
	s18 := strings.Fields("a b c")
	fmt.Println(s18)
	fmt.Println(s18[1])
}