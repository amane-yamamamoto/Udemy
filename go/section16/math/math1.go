package main

import (
	"fmt"
	"math"
)

func main(){
	// 円周率
	fmt.Println(math.Pi)

	// ２の正の平方根
	fmt.Println(math.Sqrt2)

	// float32で表現可能な最大値
	fmt.Println(math.MaxFloat32)
	// float32で表現可能な0ではない最小値
	fmt.Println(math.SmallestNonzeroFloat32)
	// float64で表現可能な最大値
	fmt.Println(math.MaxFloat64)
	// float64で表現可能な0ではない最小値
	fmt.Println(math.SmallestNonzeroFloat64)
	// int64 ver
	fmt.Println(math.MaxInt64)
	fmt.Println(math.MinInt64)
}