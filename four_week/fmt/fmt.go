package main

import (
	"fmt"
	"os"
	"bufio"
)

type Student struct{
	Id int
	name string
}

func main ()  {
	var cmd string
	var name string
	var id int
	var line string
	var s Student
	f := bufio.NewReader(os.Stdin)
	
	for {
		fmt.Print(">")
		line, _ = f.ReadString('\n')
		fmt.Sscan(line, &cmd)
		switch cmd {
		case "list":
			fmt.Println("list")
		case "add":
			fmt.Sscan(line, &cmd,&name, &id)
			s.Id = id
			s.name = name
			fmt.Println(name, id)
			
		}
	}
}