package backlog

import (
	"encoding/json"
	"testing"
)

func TestParseSpace(t *testing.T) {
	rawJSON := []byte(`{
    "spaceKey": "nulab",
    "name": "Nulab Inc.",
    "ownerId": 1,
    "lang": "ja",
    "timezone": "Asia/Tokyo",
    "reportSendTime": "08:00:00",
    "textFormattingRule": "markdown",
    "created": "2008-07-06T15:00:00Z",
    "updated": "2013-06-18T07:55:37Z"
}`)

	var space Space
	if err := json.Unmarshal(rawJSON, &space); err != nil {
		t.Fatal(err)
	}

	assertString(t, "SpaceKey", space.SpaceKey, "nulab")
	assertString(t, "Name", space.Name, "Nulab Inc.")
	assertInt(t, "OwnerID", space.OwnerID, 1)
	assertString(t, "Lang", space.Lang, "ja")
	assertString(t, "Timezone", space.Timezone, "Asia/Tokyo")
	assertString(t, "ReportSendTime", space.ReportSendTime, "08:00:00")
	assertString(t, "TextFormattingRule", space.TextFormattingRule, "markdown")
	assertInt(t, "Created#Year", space.Created.Year(), 2008)
	assertInt(t, "Created#Month", int(space.Created.Month()), 7)
	assertInt(t, "Created#Day", space.Created.Day(), 6)
	assertInt(t, "Updated#Year", space.Updated.Year(), 2013)
	assertInt(t, "Updated#Month", int(space.Updated.Month()), 6)
	assertInt(t, "Updated#Day", space.Updated.Day(), 18)
}
