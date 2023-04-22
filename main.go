package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func sum(min, max int, c chan int) {
	sum := 0
	for i := min; i <= max; i++ {
		sum += i
	}
	c <- sum // send sum to c
}

func summ(min, max int) int {
	sum := 0
	for i := min; i <= max; i++ {
		sum += i
	}
	return sum
}

func main() {
	max,_ := strconv.Atoi(os.Args[1])
	// _, _ = fmt.Scanln(&max)
	min := 1

	t := time.Now()
	ans := summ(min, max)
	d := time.Now().Sub(t)
	fmt.Printf("without goroutines\nanswer: %v\ntime taken: %v\n\n", ans, d.Seconds())

	t = time.Now()

	c := make(chan int)
	go sum(min, max/2, c)
	// fmt.Println("-------------max/2: ",max/2)
	go sum(max/2+1, max, c)
	x, y := <-c, <-c // receive from c

	ans = x+y
	d = time.Now().Sub(t)

	fmt.Printf("with goroutines\nanswer: %v\ntime taken: %v\n\n", ans, d.Seconds())
}