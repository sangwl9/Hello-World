package main

import (
	"Hello-World/calculator"
	"fmt"
)

func main() {

	fmt.Println(calculator.Multyply(2, 2))
	fmt.Println(calculator.Multyply(2, 2))

	fmt.Println(calculator.Multyply(500, 200))
	fmt.Println(calculator.Multyply(1726, 8888))

	fmt.Println(calculator.Multyply(1500, 6425))
	fmt.Println(calculator.Multyply(2000, 8957))

	fmt.Println(calculator.LenAndUpper("Sangwoo Lee"))
	fmt.Println(calculator.LenAndUpper("w Lee"))
	fmt.Println(calculator.LenAndUpper("k Lee"))

	fmt.Println(calculator.LenAndUpper("Sangwoo Lee"))
	fmt.Println(calculator.LenAndUpper("w Lee"))
	fmt.Println(calculator.LenAndUpper("k Lee"))

	calculator.RepeatMe("aaa", "bbbb", "cccc", "dddd")
	calculator.RepeatMe("aaa", "bbbb", "cccc", "dddd")

	calculator.RepeatMe("kkkkk", ":::::", "wwww", "qqqq")
	calculator.RepeatMe("nnnn", "bdddbbb", "ccdddcc", "ddd")
}
