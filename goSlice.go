package main

import (
	"fmt"
	"sort"
)

func sliceCreate() {
	var slice []int

	if len(slice) == 0 {
		fmt.Println("slice is empty", slice)
	}

	slice[1] = 10
	fmt.Println(slice)
}

func sliceEach() {
	var slice = []int{1, 2, 3}
	for i := 0; i < len(slice); i++ {
		slice[i] += 10

		fmt.Println(slice[i])
	}

	for i, v := range slice {
		slice[i] = v * 2
		fmt.Println(slice[i])
	}

}

func sliceAppend() {
	var slice = []int{1, 2, 3}

	slice2 := append(slice, 4)

	printArray(slice2)

	slice2 = append(slice2, 5, 6, 7, 8)

	printArray(slice2)
}

func printArray(array []int) {
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}
}

func difSliceArray() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4, 5}

	changeArray(array)
	changeSlice(slice)

	fmt.Println("array: ", array)
	fmt.Println("slice", slice)
}

func changeArray(array2 [5]int) {
	array2[2] = 200
}

func changeSlice(slice2 []int) {
	slice2[2] = 200
}

func slicingFunc() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[1:2]

	fmt.Println("array: ", array)
	fmt.Println("slice: ", slice, len(slice), cap(slice))

	array[1] = 100

	fmt.Println("After change second element")
	fmt.Println("array: ", array)
	fmt.Println("slice: ", slice, len(slice), cap(slice))

	slice = append(slice, 500)

	fmt.Println("After append 500")
	fmt.Println("array: ", array)
	fmt.Println("slice: ", slice, len(slice), cap(slice))
}

func sliceCopy() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, len(slice1))

	for i, v := range slice1 {
		slice2[i] = v
	}

	slice1[1] = 100
	fmt.Println(slice1)
	fmt.Println(slice2)
}

func sliceRemove() {
	slice := []int{1, 2, 3, 4, 5, 6}
	idx := 2

	for i := idx + 1; i < len(slice); i++ {
		slice[i-1] = slice[i]
	}

	slice = slice[:len(slice)-1]
	fmt.Println(slice)
}

func sliceAdd() {
	slice := []int{1, 2, 3, 4, 5, 6}
	slice = append(slice, 0)

	idx := 2

	for i := len(slice) - 2; i >= idx; i-- {
		slice[i+1] = slice[i]
	}

	slice[idx] = 100

	fmt.Println(slice)

}

func sliceSort() {
	s := []int{5, 2, 6, 3, 1, 4}
	sort.Ints(s)
	fmt.Println(s)
}

type StudentStruct struct {
	Name string
	Age  int
}

type Students []StudentStruct

func (s Students) Len() int {
	return len(s)
}

func (s Students) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func sliceStructSort() {
	s := []StudentStruct{
		{"화랑", 31},
		{"백두산", 52},
		{"류", 42},
		{"켄", 38},
		{"송하나", 18},
	}

	sort.Sort(Students(s))
	fmt.Println(s)
}

func sliceEx() {
	//sliceCreate()
	//sliceEach()
	//sliceAppend()
	//difSliceArray()
	//slicingFunc()
	//sliceCopy()
	//sliceRemove()
	//sliceAdd()
	//sliceSort()
	sliceStructSort()
}
