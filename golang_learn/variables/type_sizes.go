package main

import "fmt"

// Our Textio customers want to know how long they have had accounts with us.

// create a accountAgeInt variable and assign it the value of accountAgeFloat truncated to an integer.
func main() {
	accountAgeFloat := 2.6
	accountAgeInt := int64(accountAgeFloat)

	fmt.Println("Your account has existed for", accountAgeInt, "years")
}