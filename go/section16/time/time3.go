package main

import (
	"fmt"
	"time"
)

func main() {
	// 時刻の差分を取得
	t5 := time.Date(2021, 7, 24, 0, 0, 0, 0, time.Local)
	t6 := time.Now()
	fmt.Println(t6)

	// t5 - t6
	d2 := t5.Sub(t6)
	fmt.Println(d2)

	// 時刻を比較する
	fmt.Println(t6.Before(t5))
	fmt.Println(t6.After(t5))
	fmt.Println(t5.Before(t6))
	fmt.Println(t5.After(t6))
}