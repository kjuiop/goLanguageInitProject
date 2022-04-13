package main

import (
	"fmt"
	"goproject/hello/custompkg"

	"github.com/guptarohit/asciigraph"
	"github.com/tuckersGo/musthaveGo/ch16/expkg"
)

func packageEx() {
	custompkg.PrintCustm()
	expkg.PrintSample()

	data := []float64{3, 4, 5, 6, 9, 7, 5, 8, 5, 10, 2, 7, 2, 5, 6}
	graph := asciigraph.Plot(data)
	fmt.Println(graph)
}
