package backlog

import (
	"encoding/json"
	"testing"
)

func TestParseProject(t *testing.T) {
	rawJSON := []byte(`{
	"id": 1,
	"projectKey": "TEST",
	"name": "test",
	"chartEnabled": true,
	"subtaskingEnabled": true,
	"projectLeaderCanEditProjectLeader": true,
	"textFormattingRule": "markdown",
	"archived": true
}`)
	var project Project
	if err := json.Unmarshal(rawJSON, &project); err != nil {
		t.Fatal(err)
	}

	assertInt(t, "ID", project.ID, 1)
	assertString(t, "ProjectKey", project.ProjectKey, "TEST")
	assertString(t, "Name", project.Name, "test")
	assertBool(t, "ChartEnabled", project.ChartEnabled, true)
	assertBool(t, "SubtaskingEnabled", project.SubtaskingEnabled, true)
	assertBool(t, "ProjectLeaderCandEditProjectLeader", project.ProjectLeaderCanEditProjectLeader, true)
	assertString(t, "TextFormattingRule", project.TextFormattingRule, "markdown")
	assertBool(t, "Archived", project.Archived, true)
}
