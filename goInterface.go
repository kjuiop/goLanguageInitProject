package main

import "fmt"

func PrintVal(v interface{}) {
	switch t := v.(type) {
	case int:
		fmt.Printf("v is int %d\n", int(t))
	case float64:
		fmt.Printf("v is float64 %f\n", float64(t))
	case string:
		fmt.Printf("v is string %s\n", string(t))
	default:
		fmt.Printf("Not support type: %T:%v\n", t, t)
	}
}

type Ace struct {
	Age int
}

type Stringer interface {
	String() string
}

type Child struct {
	Name string
	Age  int
}

func (s Child) String() string {
	return fmt.Sprintf("안녕! 나는 %d살 %s라고 해", s.Age, s.Name)
}

func PrintChildAge(stringer Stringer) {
	s := stringer.(*Child)
	fmt.Printf("Age : %d\n", s.Age)
}

func childInterface() {
	child := Child{"철수", 12}
	var stringer Stringer

	stringer = child

	fmt.Printf("%s\n", stringer.String())
}

func PrintValueByType() {
	PrintVal(10)
	PrintVal(3.14)
	PrintVal("Hello")
	PrintVal(Ace{12})
}

func interfaceEx() {
	//childInterface()
	//PrintValueByType()
}
