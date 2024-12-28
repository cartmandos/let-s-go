package hello

import (
	"fmt"
	"strings"
)

// Say takes a string name and returns a greeting string
func Say(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

// SayV2 takes a slice of names and returns a greeting string.
func SayV2(names []string) string {
	if len(names) == 0 {
		names = []string{"world"}
	}

	return "Hello, " + strings.Join(names, ", ") + "!"
}