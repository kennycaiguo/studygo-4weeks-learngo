package main

import (
	"fmt"
	"bufio"
	"os"
)

type Student struct {
	Id int
	Name string
}

func main()  {
	var name string
	var id int
	var cmd string
	var line string
	s := make(map[string]Student)	
	f := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">")
		line, _ = f.ReadString('\n')
		fmt.Sscan(line, &cmd)
		switch cmd {
		case "list":
			fmt.Sscan(line, &cmd, &name)
			fmt.Printf("id: %d,name: %s\n",s[name].Id,s[name].Name)
		case "add":
			fmt.Sscan(line, &cmd, &name, &id)
			s[name] = Student{
				Id: id,
				Name: name,
			}
			
		}
		
	}
}