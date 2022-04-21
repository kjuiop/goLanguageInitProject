package main

import "fmt"

type StringerA interface {
	String() string
}

type StudentC struct {
	Age int
}

func (s *StudentC) String() string {
	return fmt.Sprintf("Student Age: %d", s.Age)
}

func PrintAge(stringer StringerA) {
	s := stringer.(*StudentC)
	fmt.Printf("Age: %d\n", s.Age)
}

func includeInterfaceEx() {
	s := &StudentC{15}
	PrintAge(s)
}
