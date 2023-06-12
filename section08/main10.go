package main

import "fmt"

func init() {
	fmt.Println("init")
}

// 複数定義可能だが、分けるメリットはない
func init() {
	fmt.Println("init2")
}

func main() {
	fmt.Println("Main")
}