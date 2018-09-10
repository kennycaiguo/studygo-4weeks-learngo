package main

import (
	"fmt"
	"unsafe"
)

func main()  {
	a1 := [3]int{1, 2, 3}
	var a2 [3]int
	a2 = a1
	fmt.Println(&a1[0], &a2[0])
	fmt.Println(unsafe.Sizeof(a1))
}