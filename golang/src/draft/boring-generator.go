package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	c := boring("boring!")
	for i:=0; i<5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You are boring, I am leaving.")
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i:=0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	// return the channel to the caller
	return c
}
