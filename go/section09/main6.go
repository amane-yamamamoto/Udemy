package main

import "fmt"

// map
func main() {

	// 明示的
	var m = map[string]int{"A": 100, "B": 200} 
	fmt.Println(m)

	// 暗黙的　
	m2 := map[string]int{"A": 100, "B": 200}
	fmt.Println(m2)

	// make
	m4 := make(map[int]string)
	fmt.Println(m4)

	m4[1] = "JAPAN"
	m4[2] = "USA"
	fmt.Println(m4)

	fmt.Println(m["A"])
	fmt.Println(m4[2])
	fmt.Println(m4[3])  // 初期値を取り出せる

	// 存在しない場合のためのエラーハンドリング
	s, ok := m4[3]
	if !ok {
		fmt.Println("error")
	}
	fmt.Println(s, ok)

	// 更新
	m4[2] = "US"
	fmt.Println(m4)
	m4[3] = "CHINA"
	fmt.Println(m4)

	// 削除 delete(map, key)
	delete(m4, 3)
	fmt.Println(m4)

	fmt.Println(len(m4))
}