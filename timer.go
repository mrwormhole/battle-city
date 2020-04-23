package main

import "time"

type Timer struct {
	time      time.Time
	deltaTime float64
}

func createTimer() *Timer {
	return &Timer{time: time.Now()}
}

func (timer *Timer) setTime() {
	timer.time = time.Now()
}

func (timer *Timer) setDeltaTime() {
	timer.deltaTime = time.Since(timer.time).Seconds() * TARGET_FPS
}
