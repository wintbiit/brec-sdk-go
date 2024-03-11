package webhook_test

import (
	"testing"

	"github.com/wintbiit/brec-sdk-go/webhook"
)

func TestWebhook(t *testing.T) {
	wh := webhook.NewWebhookServer("test")
	wh.OnRecordStart(func(event *webhook.EventRecordStart) {
		t.Log(event)
	})

	webhook.StartServers(":8080")
}
