package main

import (
	"fmt"
	"time"
)

func main(){
	// 指定時間のスリープ
	time.Sleep(5 * time.Second)
	fmt.Println("５秒停止後表示")
}