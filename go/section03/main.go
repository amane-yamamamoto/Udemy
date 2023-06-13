package main

import "fmt"

// i5 := 500
var i5 int = 500;

func outer() {
	var s4 string = "outer"
	fmt.Println(s4)
}

func main() {

	// 明示的な定義
	var i int = 10
	fmt.Println(i)

	var s string = "Hello go"
	fmt.Println(s)
	
	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int = 200
		s2 string = "Golang"
	)
	fmt.Println(i2, s2)

	// without initialization
	var i3 int
	var s3 string
	fmt.Println(i3, s3)

	i3 = 300
	s3 = "Go"
	fmt.Println(i3, s3)

	// Update
	i = 150;
	fmt.Println(i)

	// 暗黙的な定義
	i4 := 400
	fmt.Println(i4)

	i4 = 450
	fmt.Println(i4)

	// 再定義不可
	// i4 := 500
	// fmt.Println(i4)

	// 別の型は不可
	// i4 = "Hello"
	// fmt.Println(i4)

	fmt.Println(i5)

	outer()
	// outer内で定義されている変数は使用不可
	// fmt.Println(s4)

	// 宣言した値は使わなければならない
	var s5 string = "Not Use"
	fmt.Println(s5)

}