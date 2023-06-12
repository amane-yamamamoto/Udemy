package main

import "fmt"

func main() {
	point := 0

	// 条件付きfor
	for point < 10 {
		fmt.Println(point)
		point++
	}

	// 古典的for
	for i := 0; i < 10; i++ {
		if i == 3 {
			continue
		}

		if i == 6 {
			break
		}
		fmt.Println(i)
	}

	// 配列でfor
	arr := [3]int{1, 2, 3}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// 範囲式for
	for i, v := range arr {
		fmt.Println(i, v)
	}

	// slice
	sl := []string{"python", "PHP", "Go"}
	for i, v := range sl {
		fmt.Println(i, v)
	}

	// map
	m := map[string]int{"apple": 100, "banana":200}
	for k, v := range m {
		fmt.Println(k, v)
	} 
}