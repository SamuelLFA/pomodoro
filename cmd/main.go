package main

import (
	"fmt"
	"time"
)

func tickFeedback(isTimerStopped *bool) {
	fmt.Println()
	for range time.Tick(time.Second) {
		if *isTimerStopped {
			break
		}
		fmt.Print("-")
	}
}

func main() {
	var minutesToWork int
	var minutesToRest int

	var isWorkTimerStopped bool
	var isRestTimerStopped bool
	var input int
	for {
		fmt.Print("\nEnter the number of minutes to work: ")
		for input <= 0 {
			fmt.Scanf("%d\r", &input)
		}
		minutesToWork = input
		input = 0

		fmt.Print("\nEnter the number of minutes to rest: ")
		for input <= 0 {
			fmt.Scanf("%d\r", &input)
		}
		minutesToRest = input
		input = 0

		fmt.Printf("\n\nPRESS ENTER TO START WORK - %d minutes", minutesToWork)
		fmt.Scanln()
		workTimer := time.NewTimer(time.Duration(minutesToWork) * time.Minute)
		isWorkTimerStopped = false
		go tickFeedback(&isWorkTimerStopped)
		<-workTimer.C
		isWorkTimerStopped = true

		fmt.Printf("\n\nPRESS ENTER TO START REST - %d minutes", minutesToRest)
		fmt.Scanln()
		restTimer := time.NewTimer(time.Duration(minutesToRest) * time.Minute)
		isRestTimerStopped = false
		go tickFeedback(&isRestTimerStopped)
		<-restTimer.C
		isRestTimerStopped = true

		fmt.Println("\n\nDONE")
	}
}
