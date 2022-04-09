package main

import "fmt"

func Add(a int, b int) int {
	return a + b
}

func PrintAvgScore(name string, math int, eng int, history int) {
	total := math + eng + history
	avg := total / 3
	fmt.Print(name, " 님 평균 점수는", avg, "입니다.")
}

func funcSubMain() {
	PrintAvgScore("김일동", 80, 73, 54)
}

func Divide(a, b int) (result int, success bool) {
	if b == 0 {
		result = 0
		success = false
		return
	}
	result = a / b
	success = true
	return
}

func printNo(n int) {
	if n == 0 {
		return
	}
	fmt.Println(n)
	printNo(n - 1)
	fmt.Println("After", n)
}

func pabonachi(n int) int {
	if n < 2 {
		return n
	}
	return pabonachi(n-2) + pabonachi(n-1)
}

func funcMain() {
	//c := Add(3, 6)
	//fmt.Println(c)

	//c, success := Divide(9, 3)
	//fmt.Println(c, success)

	//d, success := Divide(9, 0)
	//fmt.Println(d, success)

	//printNo(3)

	fmt.Println(pabonachi(9))
}
