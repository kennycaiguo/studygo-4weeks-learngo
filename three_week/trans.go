package main

import (
	"fmt"
	"strconv"
	"time"
	"math/rand"
)

func main()  {
	var n int
	var f float32
	n = 10
	f = float32(n) / 3
	fmt.Println(f * 3)
	//f = float32(n * 10)
	fmt.Println(f, n)	
	var s string
	s = strconv.Itoa(n)
	n1, err := strconv.Atoi(s)
	fmt.Println(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n1)
	var x int64
	rand.Seed(time.Now().Unix())
	x = rand.Int63()
	s = strconv.FormatInt(x, 10)
	fmt.Println(s)

}