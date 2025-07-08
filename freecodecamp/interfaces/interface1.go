package main

import (
	"fmt"
	"time"
)

func sendMessage(msg message) {
	fmt.Println(msg.getMessage())
}

type message interface {
	getMessage() string
}

// don't edit below this line

type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
}

func mainN() {
	report := sendingReport{
		reportName:    "Report1",
		numberOfSends: 5,
	}
	bm := birthdayMessage{
		birthdayTime:  time.Date(1934, 05, 01, 0, 0, 0, 0, time.UTC),
		recipientName: "MAAAAn",
	}
	sendMessage(report)
	sendMessage(bm)
}
