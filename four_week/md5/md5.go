package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
	"crypto/md5"
)

func main()  {
	var str string
	filename := os.Args[1]
	f, _ := os.Open(filename)
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		str += line
	}
	data := []byte(str)
	md5sum := md5.Sum(data)
	fmt.Printf("%x\n", md5sum)
	
}