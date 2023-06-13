package main

import "fmt"

type User struct {
	Name string
	Age int
}

// メソッド
func (u User) SayName() {
	fmt.Println(u.Name)
}

// 値渡し
func (u User) SetName(name string){
	u.Name = name
}

// 参照渡し（基本はこっちにしておく）
func (u *User) SetName2(name string){
	u.Name = name
}


func main() {
	user1 := User{Name: "user1"}
	user1.SayName()

	// 更新されない
	user1.SetName("A")
	user1.SayName()

	//　更新される
	user1.SetName2("A")
	user1.SayName()

	// ポインタで宣言しても同じ
	user2 := &User{Name: "user2"}
	user2.SetName2("B")
	user2.SayName()
}