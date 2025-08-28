package main

import "fmt"

// Textio users reported that we're billing them for wildly inaccurate amounts. They're supposed to be billed .02 dollars (2 cents) for each text message sent.
func main() {
	numMessagesFromDorls := 72
	costPerMessage := .02
	totalCost := costPerMessage * float64(numMessagesFromDorls)

	fmt.Printf("Doris spent %.2f on text messages today\n", totalCost)
}