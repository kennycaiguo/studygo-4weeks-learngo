package main

import "fmt"

func main()  {
	primes := [6]int{1,2,3,4,5,6}
	q := []int{2, 3, 5, 7, 11, 13}
	r := []bool{true, false, true, false, true}
	var s[]int = primes[1:4]
	fmt.Println(s)
	fmt.Println(&s[0])
	fmt.Println(&primes[1])
	var s1 []int
	s1 = s
	fmt.Println(&s1[0] == &s[0])
	fmt.Println(q)
	fmt.Println(r)
	//q = q[1:4]
	//fmt.Println(q)
	//q = q[:2]
	//fmt.Println(q)
	//q = q[1:]
	//fmt.Println(q)
	q = q[:0]
	printSlice(q)
	q = q[:4]
	printSlice(q)
	q = q[2:]
	printSlice(q)

}

func printSlice(s []int)  {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}