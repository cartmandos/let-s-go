package main

import (
	"bufio"
	"fmt"
	"os"
)

// cmd: go run . < test.txt
func main() {
	scan := bufio.NewScanner(os.Stdin)
	words := make(map[string]int) // initialize map, non-nil but empty

	scan.Split(bufio.ScanWords) // scan word-by-word

	for scan.Scan() {
		words[scan.Text()]++
	}

	fmt.Println(len(words), "unique words")
}
