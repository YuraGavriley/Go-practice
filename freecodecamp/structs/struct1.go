// Complete the definition of the messageToSend struct.
// It needs two fields:
// phoneNumber - an integer
// message - a string.

package main

import "fmt"

func NotMain() {
	type messageToSend struct {
		phoneNumber int
		message     string
	}
	msg := messageToSend{phoneNumber: 1234567890, message: "Sup!"}
	fmt.Println(msg)
}
