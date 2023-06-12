package main

import "fmt"

func Plus(x, y int) int {
	return x + y
}

// 複数の返り値
func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

// 返り値の変数名も指定
func Double(price int) (result int) {
	result = price * 2;
	return
}

func Noreturn() {
	fmt.Println("No Return")
	return
}

func main() {
	i := Plus(1, 2)
	fmt.Println(i)

	i2, _ := Div(9, 4)
	fmt.Println(i2)

	i4 := Double(1000)
	fmt.Println(i4)

	Noreturn()
}