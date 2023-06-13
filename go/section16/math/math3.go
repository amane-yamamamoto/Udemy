package main

import (
	"fmt"
	"math"
)

func main() {
	// math.Trunc　数値の正負を問わず、小数点以下を単純に切り捨てる
	fmt.Println(math.Trunc(3.14))
	fmt.Println(math.Trunc(-3.14))

	// math.Floor 引数の整数より小さい最大の整数を表す
	fmt.Println(math.Floor(3.5))
	fmt.Println(math.Floor(-3.5))

	// math.Ceil 引数の整数より大きい最小の整数を表す
	fmt.Println(math.Ceil(3.5))
	fmt.Println(math.Ceil(-3.5))

}
