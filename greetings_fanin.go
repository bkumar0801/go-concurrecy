package main

import (
	"fmt"
	"time"
)

func Greet(msg string, s time.Duration) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			//time.Sleep(time.Duration(rand.Intn(1e3))*time.Millisecond)
			time.Sleep(time.Second * s)
		}
	}()
	return c
}

func FanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

func main() {
	c := FanIn(Greet("Adam", 2), Greet("Eve", 1))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You are boring: I am leaving!!")
}
