package main

import "fmt"

func main()  {
	 var s []int
	 printSlice(s)
	 s = append(s, 0)
	 printSlice(s)
	 s = append(s, 1)
	 printSlice(s)
	 s = append(s, 2, 3, 4)
	 s1 := []int{5,6,7}
	 printSlice(s)
	 s = append(s,s1...)
	 printSlice(s)
	 a := make([]int, 5)
	 printSlice(a)
	 b := make([]int, 0, 5)
	 printSlice(b)
}


func printSlice(s []int)  {
	fmt.Printf("len=%d, cap=%d %v\n", len(s), cap(s), s)
}