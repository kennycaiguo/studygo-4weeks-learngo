package main

import (
	"io"
	"fmt"
	"os"
	"bufio"
	"flag"
	"log"
)

var sep = flag.Int("n", 10, "10hang")

func main()  {
	flag.Parse()
	var n1 int
	var filename string
	if os.Args[1] != "-n" {
		filename = os.Args[1]
	} else{
    filename = os.Args[3]
	}
	
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		n1 = n1 + 1
		if n1 > *sep {
			break
		}
		fmt.Print(line)
	}
	f.Close()

}