package main

import (
	"bufio"
	"fmt"
	"os"
)

func scan() {
	var a int
	var b int

	n, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}
}

func scanf() {
	var a int
	var b int
	var c int
	n, err := fmt.Scanf("%d %d %d\n", &a, &b, &c)
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b, c)
	}
}

func scanln() {
	var a int
	var b int

	n, err := fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}

}

func scanError() {
	stdin := bufio.NewReader(os.Stdin)

	var a int
	var b int

	n, err := fmt.Scanln(&a, &b)

	if err != nil {
		fmt.Println(err)
		stdin.ReadString('\n')
	} else {
		fmt.Println(n, a, b)
	}

	n, err = fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n, a, b)
	}

}
