package main

import (
	"fmt"
	"time"
)

// channel close
func receiver(name string, ch chan int) {
	for {
		i, ok := <- ch
		if !ok {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name, "END")
}

func main() {
	ch1 := make(chan int, 2)

	/*
	ch1 <- 1
	// 閉じる
	close(ch1)

	// 閉じた後は、そのchannelにデータを送ることはできないが、channelからデータを受け取ることは可能
	// ch1 <- 1    // unable
	// fmt.Println(<-ch1) // able

	// ok: channelのバッファ内が空かつcloseのときにfalse
	i, ok := <- ch1
	fmt.Println(i, ok)

	i2, ok := <- ch1
	fmt.Println(i2, ok)
	*/

	go receiver("goroutine1", ch1)
	go receiver("goroutine2", ch1)
	go receiver("goroutine3", ch1)

	i := 0
	for i < 100 {
		ch1 <- i
		i++
	}
	close(ch1)
	time.Sleep(3 * time.Millisecond)

}