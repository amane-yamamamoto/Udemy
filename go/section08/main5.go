package main

import "fmt"


// 型switch
func anything(a interface{}) {
	// fmt.Println(a)
	switch v := a.(type) {
	case string:
		fmt.Println(v + "!?")
	case int:
		fmt.Println(v + 10000)
	}
}

func main(){
	anything("aaa")
	anything(1)

	var x interface{} = 3

	//  interface -> int
	i, isInt := x.(int)
	fmt.Println(i + 2)
	fmt.Println(i, isInt)
	// fmt.Println(x + 2)

	// interface -> float64
	f, isFloat64 := x.(float64)
	fmt.Println(f, isFloat64)
	
	// 変換のときの条件分岐
	if x == nil {
		fmt.Println("None")
	}else if i, isInt := x.(int); isInt {
		fmt.Println(i, "x is Int")
	}else if s, isString := x.(string); isString {
		fmt.Println(s, "x is String")
	}else{
		fmt.Println("I don't Know")
	}

	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("I don't Know")
	}

	// 値を使いたい時
	switch v := x.(type) {
	case int:
		fmt.Println(v, "int")
	case string:
		fmt.Println(v, "string")
	default:
		fmt.Println(v, "I don't Know")
	}
}