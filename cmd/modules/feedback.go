package modules

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
