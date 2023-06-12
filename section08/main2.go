package main

import (
	"fmt"
	"strconv"
)

// エラーハンドリング
func main() {
	var s string = "10"  // a

	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("i = %T\n", i)
}