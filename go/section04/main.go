package main

import (
	"fmt"
	"strconv"
)

func main() {

	// int
	var i int = 100

	var i2 int64 = 200

	fmt.Println(i + 50)

	// 型が異なるため演算不可
	// fmt.Println(i + i2) 

	fmt.Printf("%T\n", i2)
	fmt.Printf("%T\n", int32(i2))
	fmt.Println(i + int(i2))

	// float64
	var fl64 float64 = 2.4
	fmt.Println(fl64)

	fl := 3.2
	fmt.Println(fl64 + fl)
	fmt.Printf("%T, %T\n", fl64, fl)

	var fl32 float32 = 1.2
	fmt.Printf("%T\n", fl32)
	fmt.Printf("%T\n", float64(fl32))

	zero := 0.0
	pinf := 1.0 / zero
	fmt.Println(pinf)

	ninf := -1.0 / zero
	fmt.Println(ninf)

	nan := zero / zero
	fmt.Println(nan)

	// bool型
	var t, f bool = true, false
	fmt.Println(t, f)

	// string型
	var s string = "Hello Golang"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	var si string = "300"
	fmt.Println(si)
	fmt.Printf("%T\n", si)

	// 複数行
	fmt.Printf(`test
	test
		test
	`)

	// エスケープ
	fmt.Println("\"")
	fmt.Println(`"`)

	// byte型
	fmt.Println(s[0])
	fmt.Println(string(s[0]))

	// byte型
	byteA := []byte{72, 73}
	fmt.Println(byteA)

	fmt.Println(string(byteA))

	c := []byte("HI")
	fmt.Println(c)

	fmt.Println(string(c))

	// 配列型
	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1)

	// 宣言時に初期化
	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2)

	// 暗黙的な宣言
	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	// 要素数未指定
	arr4 := [...]string{"C", "D"}
	fmt.Println(arr4)
	fmt.Printf("%T\n", arr4)

	fmt.Println(arr2[0])
	fmt.Println(arr2[1])
	fmt.Println(arr2[2])

	arr2[2] = "C"
	fmt.Println(arr2)

	// var arr5 [4]int
	// arr5 = arr1
	// fmt.Println(arr5)

	// 要素数を調べる
	fmt.Println(len(arr1))

	// interface&nil
	var x interface{}
	fmt.Println(x)

	// あらゆる型を代入可能
	x = 1
	fmt.Println(x)
	x = 3.14
	fmt.Println(x)
	x = "A"
	fmt.Println(x)
	x = [3]int{1, 2, 3}
	fmt.Println(x)

	// 計算は不可
	// x = 2
	// fmt.Println(x + 3)

	// 型変換
	var ii int = 1
	fl644 := float64(ii)
	fmt.Println(fl644)
	fmt.Printf("ii = %T\n", ii)
	fmt.Printf("fl644 = %T\n", fl644)

	i22 := int(fl644)
	fmt.Printf("i22 = %T\n", i22)

	fl2 := 10.5
	i33 := int(fl2)
	fmt.Printf("i33 = %T\n", i33)
	fmt.Println(i33)

	var ss string = "100"
	fmt.Printf("ss = %T\n", ss)

	// str -> num
	// 引数を破棄することで、使わない形にできる（）
	// ix, _ := strconv.Atoi(ss)
	ix, err := strconv.Atoi(ss)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ix)
	fmt.Printf("%T\n", ix)

	// num -> str 
	var ii2 int = 200
	sx := strconv.Itoa(ii2)
	fmt.Println(sx)
	fmt.Printf("sx = %T\n", sx)

	var h11 string = "Hello World"
	bb := []byte(h11)
	fmt.Println(bb)

	h22 := string(bb)
	fmt.Println(h22)
}