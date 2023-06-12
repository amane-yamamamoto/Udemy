package main

import "fmt"

// field_name, type
type T struct {
	User User
}

type User struct {
	Name string
	Age int
}

func (u *User) SetName() {
	u.Name = "A"
}

func main() {
	t := T{User: User{Name: "user1", Age: 10}}

	fmt.Println(t)
	fmt.Println(t.User.Name)

	// 構造体Tのフィールド名を省略した場合、以下の直接アクセスが可能
	// fmt.Println(t.Name)

	// Userのメソッドにアクセス
	t.User.SetName()
	fmt.Println(t)
}