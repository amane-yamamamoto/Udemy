package main

import "fmt"

func main() {
	n := 1
	switch n {
	case 1, 2:
		fmt.Println("1 or 2")	
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("I dont know")
	}

	switch n := 2; n {
	case 1, 2:
		fmt.Println("1 or 2")	
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("I dont know")
	}

	// 数値列挙と計算式の混在はエラーとなる
	/*
	switch n := 2; n {
	case 1, 2:
		fmt.Println("1 or 2")	
	case 3, 4:
		fmt.Println("3 or 4")
	case n > 3 && n <  6:
		fmt.Println("3 < n < 6")
	default:
		fmt.Println("I dont know")
	}
	*/
}