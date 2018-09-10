package main

import (
	"fmt"
	"os"
	"strconv"

)

func main () {
	var s1, s3 string
	s1 = os.Args[1]
	s3 = os.Args[3]

	var n1, n2 int
	n1 ,err1 := strconv.Atoi(s1)
	n2, err2 := strconv.Atoi(s3)
	if err1 != nil && err2 != nil{
		fmt.Println(err1, err2)

	} 
	fmt.Println(n1 + n2)

}