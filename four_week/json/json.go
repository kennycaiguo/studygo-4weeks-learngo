package main

import (
	"fmt"
	"encoding/json"
	"log"
)

type Student  struct {
	Id int
	Name string
}

func main()  {
	//序列化
	a := Student{
		Id: 2,
		Name: "alice",
	}
	buf, err := json.Marshal(a)
	if err != nil {
		log.Fatal("marshal error: %s", err)
	}
	fmt.Println(string(buf))
	//反序列化
	str := `{"Id": 2, "Name":"alice"}`
	var s1 Student
	err1 := json.Unmarshal([]byte(str), &s1)
	if err1 != nil {
		log.Fatalf("unmarshal eror:%s",err1)
	}
	fmt.Println(s1)
}