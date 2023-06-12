package main

import "fmt"

func Double(i int) {
	i = i * 2
}

func Double2(i *int) {
	*i = *i * 2
}

func Double3(s []int) {
	for i, v := range s {
		s[i] = v * 2
	}
}

func main() {
	var n int = 100

	fmt.Println(n)
	fmt.Println(&n)

	// 値渡しなので変化しない
	Double(n)
	fmt.Println(n)

	// ポインタ宣言
	var p *int = &n
	fmt.Println(p)   // address
	fmt.Println(*p)  // value

	*p = 300
	fmt.Println(*p)
	n = 200

	Double2(&n)
	fmt.Println(n)

	Double2(p)
	fmt.Println(*p)

	// 元々参照渡しになるもの（参照型）
	var sl []int = []int{1, 2, 3} 
	Double3(sl)
	fmt.Println(sl)
}