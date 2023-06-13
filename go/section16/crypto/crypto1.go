package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func main() {
	// MD5のハッシュ値を生成
	// 任意の文字列からMD5のハッシュ値を生成する処理例
	h := md5.New()

	io.WriteString(h, "ABCDE")

	// ハッシュ値のバイト配列を生成
	fmt.Println(h.Sum(nil))

	// １６進数文字列を生成
	s := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(s)
}