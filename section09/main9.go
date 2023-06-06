package main

import (
	"fmt"
	"time"
)

// 受け取ったものを表示
func reciever(c chan int) {
	for {
		i := <- c
		fmt.Println(i)
	}
}

func main() {
	// // mainに送信する他のルーチンが存在しないため、dead lock
	// ch1 := make(chan int)
	// fmt.Println(<-ch1)

	ch1 := make(chan int)
	ch2 := make(chan int)
	
	go reciever(ch1)
	go reciever(ch2)

	i := 0
	for i < 100 {
		ch1 <- i
		ch2 <- i
		time.Sleep(50 * time.Millisecond)
		i++
	}

}