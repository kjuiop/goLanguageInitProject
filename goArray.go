package main

import "fmt"

func rageArray() {
	var t [5]float64 = [5]float64{24.0, 25.9, 27.8, 26.9, 26.2}

	for i, v := range t {
		fmt.Println(i, v)
	}
}

func baseArray() {
	var t [5]float64 = [5]float64{24.0, 25.9, 27.8, 26.9, 26.2}

	for i := 0; i < 5; i++ {
		fmt.Println(t[i])
	}
}

func arrayEx() {
	//baseArray()
	rageArray()
}
