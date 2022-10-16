package main

import (
	"fmt"

	"github.com/juunini/simple-go-line-notify/notify"
	"github.com/robfig/cron"
)

var accessToken = "xxxxx"

func runCronJobs() {
	s := cron.New()
	s.AddFunc("* * * * *", func() {
		message := "hello, Thepnatee! (Golang)"
		if err := notify.SendText(accessToken, message); err != nil {
			panic(err)
		}
	})
	s.Start()
}
func main() {
	runCronJobs()
	fmt.Scanln()
}
