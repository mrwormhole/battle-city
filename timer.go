package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type Timer struct {
	currentTime uint32
	lastTime    uint32
	delay       uint32
	deltaTime   float64
}

func createTimer(delay uint32) *Timer {
	return &Timer{delay: delay}
}

func (timer *Timer) start() {
	timer.currentTime = sdl.GetTicks()
}

func (timer *Timer) tick() bool {

	if timer.currentTime > timer.lastTime+timer.delay {
		fmt.Println(timer.delay)
		timer.deltaTime = float64(timer.currentTime-timer.lastTime) / 1000.0
		fmt.Println(timer.deltaTime)
		timer.lastTime = timer.currentTime
		return true
	}
	return false
}
