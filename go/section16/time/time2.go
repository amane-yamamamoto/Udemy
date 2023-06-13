package main

import (
	"fmt"
	"time"
)

func main() {
	// 時刻の感覚を表す
	fmt.Println(time.Hour)
	fmt.Printf("%T\n", time.Hour)
	fmt.Println(time.Minute)
	fmt.Println(time.Second)
	fmt.Println(time.Millisecond)
	fmt.Println(time.Microsecond)
	fmt.Println(time.Nanosecond)

	// time.Duration型を文字列から生成する
	d, _ := time.ParseDuration("2h30m")
	fmt.Println(d)

	// 現在時刻の２分３０秒後を表すtime.Time型の取得
	t3 := time.Now()
	t3 = t3.Add(2*time.Minute + 15*time.Second)	fmt.Println(t3)
}