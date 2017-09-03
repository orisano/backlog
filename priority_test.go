package backlog

import (
	"encoding/json"
	"testing"
)

func TestParsePriority(t *testing.T) {
	rawJSON := []byte(`{
	"id": 3,
	"name": "中"
}`)
	var priority Priority
	if err := json.Unmarshal(rawJSON, &priority); err != nil {
		t.Fatal(err)
	}

	assertInt(t, "ID", priority.ID, 3)
	assertString(t, "Name", priority.Name, "中")
}
