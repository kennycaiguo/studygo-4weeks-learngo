package main

import (
	"fmt"
	"os"
	"log"
	"strconv"
	"bufio"
)

func main()  {
	f, err := os.Open("/proc")
	if err != nil {
		log.Fatal(err)
	}
	infos, err := f.Readdir(-1)
	for _, info := range infos {
		if !info.IsDir() {
			continue
		}
		n, err := strconv.Atoi(info.Name())
		if err != nil {
				continue
		}
		list(n ,info.Name())
	}
	f.Close()
	
}

func list(n int,filename string)  {
	f1, _ := os.Open("/proc/" + filename + "/cmdline")
	r := bufio.NewReader(f1)
	line, _ := r.ReadString('\n')
	fmt.Printf("%d %v\n",n, line)
	f1.Close()
}