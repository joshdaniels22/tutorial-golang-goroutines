package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func summ(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	t := time.Now()
	ans := summ(s)
	d := time.Now().Sub(t)
	fmt.Printf("ans: %v -- time: %v\n",ans,d.Nanoseconds())

	t = time.Now()

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	ans = x+y
	d = time.Now().Sub(t)

	fmt.Printf("ans: %v -- time: %v\n",ans,d.Nanoseconds())
}