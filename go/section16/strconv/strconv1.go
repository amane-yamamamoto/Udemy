package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 真偽値を文字列に変換
	bt := true
	fmt.Printf("%T\n", strconv.FormatBool(bt))

	// 整数を文字列にする
	i := strconv.FormatInt(-100, 10)
	fmt.Printf("%v, %T\n", i, i)

	// 簡易的に変換できる
	i2 := strconv.Itoa(100)
	fmt.Printf("%v, %T\n", i2, i2)

	// 浮動小数点型に変換する
	fmt.Println(strconv.FormatFloat(123.456, 'E', -1, 64))
	// 浮動小数点型に変換する（小数点以下２桁まで）
	fmt.Println(strconv.FormatFloat(123.456, 'e', 2, 64))
	// 実数表現による文字列化
	fmt.Println(strconv.FormatFloat(123.456, 'f', -1, 64))
	// 浮動小数点型に変換する（小数点以下２桁まで）
	fmt.Println(strconv.FormatFloat(123.456, 'f', 2, 64))
	
	// 文字列から真偽値に変換
	bt1, _ := strconv.ParseBool("true")
	fmt.Printf("%v, %T", bt1, bt1)
	bt2, _ := strconv.ParseBool("1")
	bt3, _ := strconv.ParseBool("t")
	bt4, _ := strconv.ParseBool("T")
	bt5, _ := strconv.ParseBool("TRUE")
	bt6, _ := strconv.ParseBool("True")
	fmt.Println(bt2, bt3, bt4, bt5, bt6)


	// bf1, ok := strconv.ParseBool("false")
	// if ok != nil {
	// 	fmt.Println("Convert Error")
	// }

	// 文字列を整数に変換
	i3, _ := strconv.ParseInt("12345", 10, 0)
	fmt.Printf("%v, %T\n", i3, i3)

	// 簡易変換
	i10, _ := strconv.Atoi("123")
	fmt.Println(i10)


	// 文字列を浮動小数点に変換
	fl1, _ := ParseFloat("3.14", 64)
	fmt.Println(fl1)
}