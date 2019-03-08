package main

import "fmt"

func main() {
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"
	// messages <- "too much" // fatal error: all goroutines are asleep - deadlock!

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
