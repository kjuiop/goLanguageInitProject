package main

import "fmt"

func multiParam() {

	day := "thursday"

	switch day {
	case "monday", "tuesday":
		fmt.Println("월, 화요일은 수업을 가는 날")
	case "wednesday", "thursday", "friday":
		fmt.Println("수, 목, 금요일은 실습 가는 날입니다.")
	}
}

func fallthroughSwitch() {
	a := 3

	switch a {
	case 1:
		fmt.Println("a == 1")
	case 2:
		fmt.Println("a == 2")
	case 3:
		fmt.Println("a == 3")
		fallthrough
	case 4:
		fmt.Println("a == 4")
	default:
		fmt.Println(" a > 5")
	}
}

func switchEx() {
	//multiParam()
	fallthroughSwitch()
}
