package main

import "fmt"

func main()  {
	ages := make(map[string]int)
	ages["a"] = 1
	ages["b"] = 2

	ages1 := map[string]int{
		"a": 1,
		"b": 2,
	}
	fmt.Println(ages)
	fmt.Println(ages1)
	ages["a"] = ages["b"] + 2
	//判断是否存在
	c, ok := ages["a"]
	if ok {
		fmt.Println(c)
	} else {
		fmt.Println("not found")
	}
	if c, ok := ages["b"]; ok {
		fmt.Println(c)
	}
	//删除元素
	delete(ages, "a")
	fmt.Println(ages)

	//遍历,遍历顺序无序，不确定
	for name, age := range ages {
		fmt.Println("name", name, "age", age)
	}
	for name := range ages1 {
    fmt.Println("name", name) 
	}
}
//注意：用var申明的map是nil结构，不能进行任何操作，需要make赋初值后才能进行操作