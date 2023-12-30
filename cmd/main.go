package main

import (
	"fmt"

	"github.com/SamuelLFA/pomodoro/cmd/modules"
)

func main() {
	var minutesToWork int
	var minutesToRest int

	var isWorkTimerStopped bool
	var isRestTimerStopped bool
	for {
		minutesToWork, minutesToRest = modules.GetUserSettings(minutesToWork, minutesToRest)

		fmt.Printf("\n\nPRESS ENTER TO START WORK - %d minutes", minutesToWork)
		fmt.Scanln()
		modules.RunTimer(minutesToWork, isWorkTimerStopped)

		fmt.Printf("\n\nPRESS ENTER TO START REST - %d minutes", minutesToRest)
		fmt.Scanln()
		modules.RunTimer(minutesToRest, isRestTimerStopped)

		fmt.Println("\n\nDONE")
	}
}
