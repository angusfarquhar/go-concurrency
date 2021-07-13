package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Routines
func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

// Channels
func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	// for i := 0; i < 10; i++ {
	// 	go f(i)
	// }
	// var input string
	// fmt.Scanln(&input)

	var c chan string = make(chan string)

	go pinger(c)
	go printer(c)
	go ponger(c)

	var input string
	fmt.Scanln(&input)
}
