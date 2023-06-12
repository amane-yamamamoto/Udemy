package main

import "fmt"

// channel
func main(){
	var ch1 chan int

	// for receiving
	// var ch2 <-chan int

	// for sending
	// var ch3 chan<- int

	ch1 = make(chan int)
	ch2 := make(chan int)

	fmt.Println(cap(ch1))
	fmt.Println(cap(ch2))

	// バッファサイズ５のチャネルを確保
	ch3 := make(chan int, 5)
	fmt.Println(cap(ch3))

	// send to channel
	ch3 <- 1
	fmt.Println(len(ch3))
	ch3 <- 2
	ch3 <- 3
	fmt.Println("len", len(ch3))

	// receive from channel
	i := <- ch3
	fmt.Println(i)
	fmt.Println("len", len(ch3))

	i2 := <- ch3
	fmt.Println(i2)
	fmt.Println("len", len(ch3))

	// 変数に入れずに受けとる
	fmt.Println(<-ch3)
	fmt.Println("len", len(ch3))

	// バッファサイズを超えたデータ量を送ると、dead lockになる
	// ch3 <- 1
	// ch3 <- 2
	// ch3 <- 3
	// ch3 <- 4
	// ch3 <- 5
	// ch3 <- 6
}