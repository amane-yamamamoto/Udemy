package main

import "fmt"

// 定数
const Pi = 3.14

const (
	URL = "https:xxx.co.jp"
	SiteName = "test"
)

// 省略
const (
	A = 1
	B
	C
	D = "A"
	E
)

// 数値連番
const (
	c0 = iota
	c1
	c2
)

// 最大値
// var Big int = 9223372036854775807 + 1 
const Big = 9223372036854775807 + 1 

func main() {
	fmt.Println(Pi)

	// 上書き不可
	// Pi = 3

	fmt.Println(URL, SiteName)

	fmt.Println(A, B, C, D, E)
	
	fmt.Println(Big - 1)

	fmt.Println(c0, c1, c2)
}