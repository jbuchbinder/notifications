package notifications

import (
	"testing"
)

func Test_HipChat(t *testing.T) {
	hc := &NotifierHipChat{target: "", credentials: ""}
	hc.Init()
	hc.Notify("Test message from new client")
}
