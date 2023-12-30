package modules

import (
	"time"
)

func RunTimer(minutesToWork int, isWorkTimerStopped bool) {
	workTimer := time.NewTimer(time.Duration(minutesToWork) * time.Minute)
	isWorkTimerStopped = false
	go tickFeedback(&isWorkTimerStopped)
	<-workTimer.C
	isWorkTimerStopped = true
}
