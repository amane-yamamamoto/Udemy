package main

import "fmt"

// append main len cap
func main() {
	sl := []int{100, 200}
	fmt.Println(sl)

	// 追加
	sl = append(sl, 300)
	fmt.Println(sl)

	sl = append(sl, 400, 500, 600)
	fmt.Println(sl)

	// make
	sl2 := make([]int, 5)
	fmt.Println(sl2)

	fmt.Println(len(sl2))
	fmt.Println(cap(sl2))

	// make(_, len, cap)
	sl3 := make([]int, 5, 10)
	fmt.Println(len(sl3))
	fmt.Println(cap(sl3))

	sl3 = append(sl3, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(len(sl3))
	fmt.Println(cap(sl3)) // 容量が増大する

}

