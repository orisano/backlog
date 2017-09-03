package backlog

import (
	"encoding/json"
	"testing"
)

func TestParseNotification(t *testing.T) {
	rawJSON := []byte(`{
    "content": "Notification",
    "updated": "2013-06-18T07:55:37Z"
}`)

	var notification Notification
	if err := json.Unmarshal(rawJSON, &notification); err != nil {
		t.Fatal(err)
	}

	assertString(t, "Content", notification.Content, "Notification")
	assertTime(t, "Updated", notification.Updated, 2013, 6, 18)
}
