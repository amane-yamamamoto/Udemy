package main

import "fmt"

type User struct {
	Name string
	Age  int
	// X, Y int
}

// 値渡し
func UpdateUser(user User) {
	user.Name = "A"
	user.Age = 1000
}

// 参照渡し
func UpdateUser2(user *User) {
	user.Name = "A"
	user.Age = 1000
}

func main() {
	var user1 User
	fmt.Println(user1)

	// 更新
	user1.Name = "user1" 
	user1.Age = 10
	fmt.Println(user1)

	user2 := User{}
	fmt.Println(user2)
	user2.Name = "user2"
	fmt.Println(user2)

	user3 := User{Name: "user3", Age: 30}
	fmt.Println(user3)

	// 順番は変えてはいけない
	user4 := User{"user4", 30}
	fmt.Println(user4)

	user6 := User{Name: "user6"}
	fmt.Println(user6)

	// 構造体のポインタ型を返す
	user7 := new(User)
	fmt.Println(user7)

	// user7と同じだが、こちらの宣言方法が多い
	// 関数の引数として使う場合が多い
	user8 := &User{}
	fmt.Println(user8)

	UpdateUser(user1)
	UpdateUser2(user8)

	fmt.Println(user1)
	fmt.Println(user8)
}