package main

import (
	"os"
	"fmt"
)

func main()  {
	m := 10
	n := 3
	switch os.Args[1] {
	case "+":
		fmt.Println(m + n)
	case "-":
		fmt.Println(m - n)
	case "*":
		fmt.Println(m * n)
	case "/":
		fmt.Println(m / n)
	}
}
