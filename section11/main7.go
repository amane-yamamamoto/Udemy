package main

import "fmt"

// 独自の型を作る
type MyInt int

func (mi MyInt) Print() {
	fmt.Println(mi)
}

func main() {
	var mi MyInt
	fmt.Println(mi)
	fmt.Printf("%T\n", mi)

	// 計算不可
	// i := 100
	// fmt.Println(i + mi)

	mi.Print()
}