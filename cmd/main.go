package main

import (
	"fmt"
	"time"
)

func tickFeedback(isTimerStopped bool) {
	for range time.Tick(time.Second) {
		if isTimerStopped {
			break
		}
		fmt.Print("-")
	}
}

func main() {
	for {
		fmt.Print("\nEnter the number of minutes to work (DEFAULT 15): ")
		var minutesToWork int = 15
		fmt.Scanf("%d", &minutesToWork)

		fmt.Print("\nEnter the number of minutes to rest (DEFAULT 5): ")
		var minutesToRest int = 5
		fmt.Scanf("%d", &minutesToRest)

		var isWorkTimerStopped bool = false
		var isRestTimerStopped bool = false

		fmt.Printf("\n\nPRESS ENTER TO START WORK - %d minutes", minutesToWork)
		workTimer := time.NewTimer(time.Duration(minutesToWork) * time.Minute)
		go tickFeedback(isWorkTimerStopped)
		<-workTimer.C
		isWorkTimerStopped = true

		fmt.Printf("\n\nPRESS ENTER TO START REST - %d minutes", minutesToRest)
		restTimer := time.NewTimer(time.Duration(minutesToRest) * time.Minute)
		go tickFeedback(isRestTimerStopped)
		<-restTimer.C
		isRestTimerStopped = true

		fmt.Println("\n\nDONE")
	}
}
