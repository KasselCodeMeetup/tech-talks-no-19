package hello

import "time"

func Hello(number int) {
	for {
		println("Hello from Goroutine", number)
		time.Sleep(500 * time.Millisecond)
	}
}
