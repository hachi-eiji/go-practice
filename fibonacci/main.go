package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		r := time.Duration(rand.Intn(1000)) * time.Millisecond
		fmt.Printf("sleep %s", r)
		//time.Sleep(r)
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	fmt.Printf("%v", c)
	for i := range c {
		fmt.Println(i)
	}
	return
}
