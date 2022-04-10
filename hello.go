package main

import "fmt"

func hello() {
	fmt.Println("Hello world !")

	var a int = 10
	var msg string = "Hello Variable"

	a = 20
	msg = "Good Morning"
	fmt.Println(msg, a)

	var minimumWage int = 10
	var workingHour int = 20

	var income int = minimumWage * workingHour

	fmt.Println(minimumWage, workingHour, income)
}

func main() {

	//hello()
	//goVarType()
	//goFmt()
	//scan()
	//scanf()
	//scanln()
	//scanError()
	//cal()
	//funcMain()
	//funcSubMain()
	//constEx()
	//ifEx()
	//switchEx()
	//arrayEx()
	structEx()
}
