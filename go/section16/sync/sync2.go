package main

import (
	"fmt"
	"sync"
)
  
func main() {

	// sync.WaitGroup
	wg := new(sync.WaitGroup)
	// 待ち受けるゴルーチンの数
	wg.Add(3)

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("1st Goroutine")
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("2nd Goroutine")
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("3rd Goroutine")
		}
		wg.Done()
	}()

	// ゴルーチンの Done を待ち受ける
	wg.Wait()
}