package main

import (
	"fmt"
	"time"
)

func printMessage(msg string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(msg)
		time.Sleep(time.Second)
	}
}

func main() {
	go printMessage("Hello") // Start a goroutine to execute printMessage concurrently
	printMessage("World")    // Execute printMessage in the main goroutine
	time.Sleep(5 * time.Second)
}
