package main


import (
	"fmt"
	"os"
)

// os
func main() {
	// Exit
	os.Exit(1)
	fmt.Println("start")

	// defer 文も実行されない
	defer func() {
		fmt.Println("defer")
	}()
	os.Exit(0)
}