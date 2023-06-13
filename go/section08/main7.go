package main

import (
	"fmt"
	"os"
)

// defer
func TestDefer() {
	defer fmt.Println("END")
	fmt.Println("START")
}

func RunDefer(){
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
}

func main() {
	TestDefer()

	// 無名関数
	// defer func() {
	// 	fmt.Println("1")
	// 	fmt.Println("2")
	// 	fmt.Println("3")
	// }()

	RunDefer()

	// リソースの開放処理
	file, err := os.Create("test.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	
	file.Write([]byte("Hello"))
}