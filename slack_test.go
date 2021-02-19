package notifications

import (
	"testing"
)

func Test_Slack(t *testing.T) {
	// Attemp to send test message
	hc := &NotifierSlack{target: "", credentials: ""}
	hc.Init()
	hc.Notify("Test message from notifications")
}
