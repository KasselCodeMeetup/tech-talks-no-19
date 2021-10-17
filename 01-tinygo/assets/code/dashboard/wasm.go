package main

import (
	"github.com/KasselCodeMeetup/tech-talks-no19/01-tinygo/assets/code/dashboard/views/dashboard"
	"github.com/KasselCodeMeetup/tech-talks-no19/01-tinygo/assets/code/dashboard/views/login"
)

func main() {
	dashboardService := dashboard.New()

	loginEvents := make(chan login.Event, 1)
	loginService := login.New(loginEvents)
	loginService.RenderLogin()

	loginEvent := <-loginEvents

	println("New user logged in:", loginEvent.UserName)

	dashboardService.RenderDashboard()
	select {}
}
