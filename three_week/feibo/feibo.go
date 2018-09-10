package main

import (
	"fmt"
	"os"
	"strconv"
)

func main()  {
	s := os.Args[1]
	var sum1 int
	var temp  int
	var temp1 int
	var temp2 int
	var n1 int
	var n2 int
	var n int
	n, err := strconv.Atoi(s)
	if err != nil {
		return
	}
	n1 = 1
	n2 = 1
	for {
		temp1 = n1
		temp2 = n2
		temp = temp1 + temp2  
		sum1 += temp
		n1 = temp
		fmt.Printf("temp1: %d, temp2: %d temp: %d\n", temp1, temp2, temp)
		if n1 >= n {
			fmt.Println(temp1)
			fmt.Println(sum1 - temp2)
			break
		}
		n2 = n1 + temp2
	}
}