package main

import "fmt"

func main() {

	// 算術演算子
	fmt.Println(1 + 2)
	fmt.Println("ABC" + "DE")

	fmt.Println(5 - 1)
	fmt.Println(5 * 4)
	fmt.Println(60 / 3)
	fmt.Println(9 % 4)

	n := 0
	n += 4
	fmt.Println(n)
	n++
	fmt.Println(n)
	n--
	fmt.Println(n)

	s := "ABC"
	s += "DE"
	fmt.Println(s)

	// 比較演算子
	fmt.Println(1 == 1)
	fmt.Println(1 == 2)
	fmt.Println(4 <= 8)
	fmt.Println(1 >= 8)
	fmt.Println(true != false)

	// 論理演算子
	fmt.Println(true && false == true)
	fmt.Println(true || false == true)
	fmt.Println(!true)
}