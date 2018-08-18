package main

import "fmt"

func main()  {
	var (
		x int 
		y float32
		z string
		p *int
	)
	i := 0
	s := "hello"
	i, j := 2, 1
	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", y)
	fmt.Printf("%v\n", z)
	fmt.Printf("%v\n", p)
	fmt.Printf("%v\n", i)
	fmt.Printf("%v\n", s)
	fmt.Printf("%v\n", i)
	fmt.Printf("%v\n", j)
}