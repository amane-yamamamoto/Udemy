package main

import (
	"log"
	"os"
	"fmt"
)

func main() {
	// OpenFile
	// O_RDONLY
	// O_WRONLY
	// O_RDWR
	// O_APPEND
	// O_CREATE
	// O_TRUNC
	f, err := os.OpenFile("foo.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	bs := make([]byte, 128)
	n, err := f.Read(bs)
	fmt.Println(n)
	fmt.Println(string(bs))
}