package notifications

import (
	"log"

	"github.com/tbruyelle/hipchat-go/hipchat"
)

func init() {
	//log.Printf("Discovered notification backend 'hipchat'")
	notifierMap["hipchat"] = &NotifierHipChat{}
}

type NotifierHipChat struct {
	credentials string
	debug       bool
	target      string
	//client *hipchat.Client
}

func (n NotifierHipChat) Init() {
	//n.client = hipchat.NewClient(config.Config.Notifications.HipChat.AuthToken)
}

func (n *NotifierHipChat) SetCredentials(creds []string) {
	if len(creds) > 0 {
		n.credentials = creds[0]
	}
}

func (n *NotifierHipChat) SetDebug(debug bool) {
	n.debug = debug
}

func (n *NotifierHipChat) SetTarget(target string) {
	n.target = target
}

func (n NotifierHipChat) Notify(message string) {
	if n.debug {
		log.Printf("NotifierHipChat.Notify(): [%s] %s", n.target, message)
	}

	if n.target == "" || n.target == "0" {
		log.Print("NotifierHipChat.Notify(): Ignoring null/zero room")
		return
	}

	client := hipchat.NewClient(n.credentials)
	notificationRequest := &hipchat.NotificationRequest{
		Message: message,
		Notify:  true,
	}
	_, err := client.Room.Notification(n.target, notificationRequest)
	if err != nil {
		log.Printf("NotifierHipChat.Notify(): %s", err)
	}
}
