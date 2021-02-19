package notifications

import (
	"log"
)

func init() {
	//log.Printf("Discovered notification backend 'none'")
	notifierMap["none"] = &NotifierNone{}
}

type NotifierNone struct {
	debug bool
}

func (n NotifierNone) Init() {
}

func (n *NotifierNone) SetCredentials(creds []string) {
}

func (n *NotifierNone) SetDebug(debug bool) {
	n.debug = debug
}

func (n *NotifierNone) SetTarget(target string) {
}

func (n NotifierNone) Notify(message string) {
	if n.debug {
		log.Printf("NotifierNone.Notify(): %s", message)
	}
}
