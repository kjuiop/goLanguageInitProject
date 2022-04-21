package main

import (
	"fmt"
	"os"
)

func sum(nums ...int) int {
	sum := 0

	fmt.Printf("nums 타입: %T\n", nums)
	for _, v := range nums {
		sum += v
	}

	return sum
}

func printFuncMethod() {
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(10, 20))
	fmt.Println(sum())
}

func deferFunc() {

	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Failed to create a file")
		return
	}

	defer fmt.Println("반드시 호출됩니다.")
	defer f.Close()
	defer fmt.Println("파일을 닫았습니다.")

	fmt.Println("파일에 Hello World 를 씁니다.")
	fmt.Fprintln(f, "Hello World")
}

func methodAdd(a, b int) int {
	return a + b
}

func methodMul(a, b int) int {
	return a * b
}

func getOperator(op string) func(int, int) int {
	if op == "+" {
		return methodAdd
	} else if op == "*" {
		return methodMul
	} else {
		return nil
	}
}

func printOperator() {
	var operator func(int, int) int
	operator = getOperator("*")

	var result = operator(3, 4)
	fmt.Println(result)
}

func methodFunctionEx() {
	//printFuncMethod()
	//deferFunc()
	printOperator()
}
