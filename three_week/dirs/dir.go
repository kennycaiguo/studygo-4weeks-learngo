package main

import (
	"fmt"
	"os"
	"log"
)

func main()  {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	infos, err := f.Readdir(-1)
	for _, info := range infos {
		fmt.Printf("%v %d %s\n", info.IsDir(), info.Size(), info.Name())
	}
	f.Close()
	
}