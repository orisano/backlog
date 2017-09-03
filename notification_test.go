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
	assertInt(t, "Update#Year", notification.Updated.Year(), 2013)
	assertInt(t, "Update#Month", int(notification.Updated.Month()), 6)
	assertInt(t, "Update#Day", notification.Updated.Day(), 18)
}
