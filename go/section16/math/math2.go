package main

import (
	"fmt"
	"math"
)

func main() {
	// 絶対値
	fmt.Println(math.Abs(100))
	fmt.Println(math.Abs(-100))

	// 累乗を求める
	fmt.Println(math.Pow(0, 2))
	fmt.Println(math.Pow(2, 2))

	// 平方根、立方根
	fmt.Println(math.Sqrt(2))
	fmt.Println(math.Cbrt(8))

	// 最大値、最小値
	fmt.Println(math.Max(1, 2))
	fmt.Println(math.Min(1, 2))
}