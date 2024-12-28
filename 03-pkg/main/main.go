package main

import (
	"fmt"
	"hello"
	"os"
)

func main() {
	fmt.Println(hello.SayV2(os.Args[1:]))
}
