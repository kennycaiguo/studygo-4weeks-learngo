package main

import (
	"unsafe"
	"fmt"
)

func main()  {
	var (
		x1 int
		x2 int32
		x3 int64
		x4 uint
		x5 uint16
		x6 uint32
		x7 uint64
	)
	s1 := "hello" + "world"
	fmt.Println(unsafe.Sizeof(x1))
	fmt.Println(unsafe.Sizeof(x2))
	fmt.Println(unsafe.Sizeof(x3))
	fmt.Println(unsafe.Sizeof(x4))
	fmt.Println(unsafe.Sizeof(x5))
	fmt.Println(unsafe.Sizeof(x6))
	fmt.Println(unsafe.Sizeof(x7))
	array := []byte(s1)
	fmt.Println(array)
	array[0] = 72
	s1 = string(array)
	fmt.Println(s1)
	fmt.Println('a' + ('H' - 'h'))
	fmt.Println(toupper("hahjshhuisushi"))
}


func toupper(s string) string  {
	var s1 string
	array := []byte(s)
	fmt.Println(array)
	for i := 0; i<= len(array) - 1; i++ {
		array[i] -= 32
	}
	s1 = string(array)
	return s1
}