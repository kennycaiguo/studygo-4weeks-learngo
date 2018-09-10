package main

import "fmt"

type Student struct {
	Id int
	Name string
}

func main()  {
	var s Student
	s.Id = 1
	s.Name = "jack"
	fmt.Println(s)
	s1 := Student{
		Id: 2,
		Name: "alice",
	}
	fmt.Println(s1)
	s1 = s
	fmt.Println(s1)
	var p *Student
	p = &s1
	p.Id = 3
	fmt.Println(s1)
	var p1 *int
	p1 = &s1.Id
	*p1 = 4
	fmt.Println(s1)
}