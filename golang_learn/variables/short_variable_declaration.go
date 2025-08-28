package main

import "fmt"

// A common use case for Textio is to send birthday messages.
func main() {
	messageStart := "Happy birthday! you are now"
	age := 27
	messageEnd := "years old!"
	
	fmt.Println(messageStart, age, messageEnd)
}