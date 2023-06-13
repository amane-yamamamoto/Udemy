package main

import (
	"os"
)

func main() {
	// Create
	f, _ := os.Create("foo.txt")
	f.Write([]byte("Hello\n"))
	f.WriteAt([]byte("Golang"), 6)
	f.Seek(0, os.SEEK_END)
	f.WriteString("Yaah")
}