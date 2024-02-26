package main

import gosxnotifier "github.com/deckarep/gosx-notifier"

func main() {
}

func notify(message string) error {
	notif := gosxnotifier.NewNotification(message)
	notif.Title = "Time's Up!"
	notif.Group = "me.reimarrosas.gt"
	notif.AppIcon = "icon.svg"

	return notif.Push()
}
