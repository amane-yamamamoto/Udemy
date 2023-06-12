package main

import (
	"fmt"
	f "fmt"
	"Udemy/section13/foo"
)

func appName() string {
	const AppName = "GoApp"
	var Version string = "1.0"
	return AppName + " " + Version
}

func Do(s string) (b string) {
	// var b string = s 再定義となりエラー
	// 再定義したい場合は以下（本来は避けるべき）
	b = s
	{
		b := "BBBB"
		fmt.Println(b)
	}
	fmt.Println(b)
	return b
}

func main() {
	f.Println(foo.Max)
	// fmt.Println(foo.min)

	f.Println(foo.ReturnMin())

	fmt.Println(appName())
	// fmt.Println(AppName, Version) これは当然エラー

	fmt.Println(Do("aaa"))
}