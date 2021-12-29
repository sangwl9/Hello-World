package calculator

import (
	"fmt"
	"strings"
)

// basic form
func Multyply(a int, b int) int {
	return a * b
}

// multiple return values
func LenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

// multiple parameters
func RepeatMe(words ...string) {
	fmt.Println(words)
}
