package main

import (
	"machine"
	"time"
)

func main() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	go func() {
		for {
			println("Hello from Goroutine #1 !")
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			println("Hello from Goroutine #2 !")
			time.Sleep(500 * time.Millisecond)
		}
	}()

	for {
		led.High()
		time.Sleep(time.Second)
		led.Low()
		time.Sleep(time.Second)
	}
}
