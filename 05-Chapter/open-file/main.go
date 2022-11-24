package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	f, err := os.Open("05-Chapter/open-file/file.txt")
	if err != nil {
		panic(err)
	}
	c, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	fmt.Printf("### File content ###\n%s\n", string(c))
	f.Close()
	f, err = os.OpenFile("05-Chapter/open-file/file.txt", os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}
	io.WriteString(f, "Test string")
	f.Close()

}
