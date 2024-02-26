package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	gosxnotifier "github.com/deckarep/gosx-notifier"
)

var tickCount uint

func main() {
	var timerType string
	flag.StringVar(&timerType, "type", "timer", "Timer Type: timer/interval")

	flag.Parse()

	if flag.NArg() != 1 {
		log.Fatalln("Invalid number of arguments. Must pass a duration")
	}

	timerType = strings.ToLower(timerType)
	if timerType != "timer" && timerType != "interval" {
		log.Fatalln("Invalid timer type. Must be timer or interval")
	}

	d, err := time.ParseDuration(flag.Arg(0))
	if err != nil {
		log.Fatalln("Malformed duration.")
	}

	if timerType == "timer" {
		t := time.After(d)
		select {
		case <-t:
			msg := fmt.Sprintf("%s has elapsed!", d)
			notify(msg)
		}
	} else {
		t := time.Tick(d)
		for {
			select {
			case <-t:
				tickCount++
				msg := fmt.Sprintf("%s has elapsed %s", d, countString(tickCount))
				notify(msg)
			}
		}
	}
}

func notify(message string) error {
	notif := gosxnotifier.NewNotification(message)
	notif.Title = "Time's Up!"
	notif.Group = "me.reimarrosas.gt"
	notif.AppIcon = "icon.svg"

	return notif.Push()
}

func countString(c uint) string {
	switch c {
	case 1:
		return "once"
	case 2:
		return "twice"
	default:
		return fmt.Sprintf("%d times", c)
	}
}
