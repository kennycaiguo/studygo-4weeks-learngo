package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
	"strings"
	"log"
)

func main()  {
	var value string
	temp := make(map[string]int)
	filename := os.Args[1]
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
		data := strings.Fields(line)
    for i:=0; i< len(data) - 1; i++ {
			value = data[i]
      if _, ok := temp[value]; ok {
        temp[value] += 1
			} else {
				temp[value] = 1
			}
		}
	}
	f.Close()
	for word, num := range temp{
		fmt.Println("world", word, "num", num)
	}
}