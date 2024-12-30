package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	PrintMoreThanOne(words)
}

func PrintMoreThanOne(words map[string]int) {
	type kv struct {
		key string
		val int
	}

	var ss []kv
	for k, v := range words {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].val > ss[j].val
	})

	for _, s := range ss {
		if s.val > 1 {
			fmt.Printf("\"%s\" appears %d times\n", s.key, s.val)
		}
	}
}
