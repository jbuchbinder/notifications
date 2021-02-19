package notifications

import (
	"errors"
)

var (
	notifierMap = map[string]Notifier{}

	NotifierService Notifier
)

// Notifier defines an interface for all notification services.
type Notifier interface {
	// Init performs housekeeping on startup
	Init()
	// Notify performs the actual notification
	Notify(string)
	// SetDebug allows internal debug information to be enabled or disabled. Defaults to false.
	SetDebug(bool)
	// SetCredentials allows one or more service-specific credential to be passed to a notifier. Most services will only use a single string.
	SetCredentials([]string)
	// SetTarget allows a message target to be specified.
	SetTarget(string)
}

// NotifierByName is a factory function which instantiates a notifier by its short name.
func NotifierByName(name string) (*Notifier, error) {
	var n Notifier
	if _, ok := notifierMap[name]; !ok {
		return nil, errors.New("Unable to instantiate notifier " + name)
	}
	n = notifierMap[name]
	n.Init()
	return &n, nil
}
