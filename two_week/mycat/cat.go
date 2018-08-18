package main

import (
	"fmt"
	"os"
	"io/ioutil"
)

func main()  {
	var filename string
	filename = os.Args[1]
	printFile(filename)
}

func printFile(name string)  {
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println(err)
    return
	}
	fmt.Println(string(buf))
}