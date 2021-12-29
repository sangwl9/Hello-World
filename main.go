package main

import (
	"Hello-World/calculator"
	"fmt"
)

func main() {
	fmt.Println("New world")
	fmt.Println("New world")
	fmt.Println("New world")

	fmt.Println(calculator.Multyply(2, 2))
	fmt.Println(calculator.LenAndUpper("Sangwoo Lee"))
	calculator.RepeatMe("aaa", "bbbb", "cccc", "dddd")
	calculator.RepeatMe("aaa", "bbbb", "cccc", "dddd")
}
