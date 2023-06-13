package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("表示")

	// fmt標準
	fmt.Print("Hello")
	// 改行あり
	fmt.Println("Hello!")


	// Println系 文字列間に半角スペースで区切られ、ラストに改行が入る
	fmt.Println("Hello", "world!")
	fmt.Println("Hello", "world!")

	// Printg系 フォーマット指定
	fmt.Printf("%s\n", "Hello")
	fmt.Printf("%#v\n", "Hello")

	// Sprint系 フォーマットした結果を文字列で返す。
	s := fmt.Sprint("Hello")
	s1 := fmt.Sprintf("%v\n", "Hello")
	s2 := fmt.Sprintln("Hello")

	fmt.Println(s)
	fmt.Println(s1)
	fmt.Println(s2)

	// Fprint系　書き込み先を指定
	fmt.Fprint(os.Stdout, "Hello")
	fmt.Fprintf(os.Stdout, "Hello")
	fmt.Fprintln(os.Stdout, "Hello")

	f, _ := os.Create("test1.txt")
	defer f.Close()

	fmt.Fprintln(f, "Fprint")
}
