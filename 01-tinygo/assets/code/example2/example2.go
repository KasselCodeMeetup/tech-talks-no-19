package main

import (
	"machine"
	"time"

	"github.com/KasselCodeMeetup/tech-talks-no19/01-tinygo/assets/code/example2/hello"
)

func main() {
	externalLED := machine.D2
	externalLED.Configure(machine.PinConfig{Mode: machine.PinOutput})

	go hello.Hello(1)
	go hello.Hello(2)

	for {
		externalLED.High()
		time.Sleep(time.Second)
		externalLED.Low()
		time.Sleep(time.Second)
	}
}
