package main

import "fmt"


// スライスfor
func main(){
	sl := []string{"A", "B", "C"}
	fmt.Println(sl)

	// 範囲for
	for i, v := range sl {
		fmt.Println(i, v)
	}

	// 古典的for
	for i := 0; i < len(sl); i++ {
		fmt.Println(sl[i])
	}

}