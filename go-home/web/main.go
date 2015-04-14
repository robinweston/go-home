package main

import (
	"io"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	now := time.Now().Local()
	day := now.Weekday()

	var message = ""
	if day == time.Saturday || day == time.Sunday {
		message = "It's the weekend. Let's forget about work!"
	} else if now.Hour() < 9 || now.Hour() > 18 {
		message = "You really shouldn't be at work now"
	} else {
		endOfDay := time.Date(now.Year(), now.Month(), now.Day(), 18, 0, 0, 0, time.Local)
		timeTillHome := endOfDay.Sub(now)
		message = "You get to go home in " + timeTillHome.String()
	}

	io.WriteString(w, "The current local time is "+now.String()+".\n"+message)
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}
