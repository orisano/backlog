package backlog

import (
	"encoding/json"
	"testing"
)

func TestParseIssueType(t *testing.T) {
	rawJSON := []byte(`{
	"id": 1,
	"projectId": 1,
	"name": "バグ",
	"color": "#990000",
	"displayOrder": 1
}`)
	var issueType IssueType
	if err := json.Unmarshal(rawJSON, &issueType); err != nil {
		t.Fatal(err)
	}

	assertInt(t, "ID", issueType.ID, 1)
	assertInt(t, "ProjectID", issueType.ProjectID, 1)
	assertString(t, "Name", issueType.Name, "バグ")
	assertString(t, "Color", issueType.Color, "#990000")
	assertInt(t, "DisplayOrder", issueType.DisplayOrder, 1)
}
