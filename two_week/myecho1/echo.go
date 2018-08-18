package main

import (
	"flag"
	"fmt"
	"strings"
)

var sep = flag.String("s", " ", "separator")
var sep1 = flag.Bool("n", false, "newline")

func main()  {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if *sep1 {
		fmt.Println()
	}
}