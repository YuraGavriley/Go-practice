/*
Textio has a bug, we've been sending texts
that are missing critical bits of information!
Before we send text messages in Textio,
we must check to make sure the required
fields have non-zero values.

Notice that the user struct is a nested struct within
the messageToSend struct. Both sender and recipient
are user struct types.

Complete the canSendMessage function.
It should return true only if the sender and
recipient fields each contain a name and a number.
If any of the default zero values are present,
return false instead.
*/

package main

import "fmt"

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	if mToSend.message == "" || mToSend.sender.name == "" ||
		mToSend.sender.number == 0 || mToSend.recipient.name == "" ||
		mToSend.recipient.number == 0 {
		return false
	}
	return true
}

func Notmain() {
	m := messageToSend{
		message:   "you have an event tomorrow",
		sender:    user{number: 16545550987, name: "He"},
		recipient: user{name: "Suzie Sall", number: 19035558973},
	}
	fmt.Println(canSendMessage(m))
}
