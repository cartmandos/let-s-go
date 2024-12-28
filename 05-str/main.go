package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)


func CountUnicode() {
	//NOTE: ASCII: 0-127
	s := "אlef"

	fmt.Printf("%8T %[1]v\n", s)
	// unicode chars
	fmt.Printf("%8T %[1]v\n", []rune(s))

	// actual bytes in memory
	b := []byte(s)
	fmt.Printf("%8T %[1]v %d\n", b, len(b))
}

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "not enough args")
		os.Exit(-1)
	}

	old, new := os.Args[1], os.Args[2]
	scan := bufio.NewScanner(os.Stdin)

	for scan.Scan() {
		s := strings.Split(scan.Text(), old)
		t := strings.Join(s, new)
		
		fmt.Println(t)
	}
}