package notifications

import (
	"log"

	"github.com/nlopes/slack"
)

func init() {
	//log.Printf("Discovered notification backend 'slack'")
	notifierMap["slack"] = &NotifierSlack{}
}

type NotifierSlack struct {
	credentials string
	debug       bool
	target      string
	//client *hipchat.Client
}

func (n NotifierSlack) Init() {
	//n.client = hipchat.NewClient(config.Config.Notifications.HipChat.AuthToken)
}

func (n *NotifierSlack) SetCredentials(creds []string) {
	if len(creds) > 0 {
		n.credentials = creds[0]
	}
}

func (n *NotifierSlack) SetDebug(debug bool) {
	n.debug = debug
}

func (n *NotifierSlack) SetTarget(target string) {
	n.target = target
}

func (n NotifierSlack) Notify(message string) {
	if n.debug {
		log.Printf("NotifierSlack.Notify(): [%s] %s", n.target, message)
	}

	if n.target == "" || n.target == "0" {
		log.Print("NotifierSlack.Notify(): Ignoring null/zero room")
		return
	}

	client := slack.New(n.credentials)
	_, _, err := client.PostMessage(
		n.target,
		slack.MsgOptionText(message, false),
	)
	if err != nil {
		log.Printf("NotifierSlack.Notify(): %s", err)
	}
}
