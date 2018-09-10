package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"bufio"
)

func main()  {
	f, err := os.OpenFile(os.Args[1], os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(line)
	}
	f.Close()
}