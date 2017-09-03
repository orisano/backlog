package backlog

import (
	"encoding/json"
	"testing"
)

func TestParseStar(t *testing.T) {
	rawJSON := []byte(`{
	"id":75,
	"comment":null,
	"url": "https://xx.backlog.jp/view/BLG-1",
	"title": "[BLG-1] first issue | 課題の表示 - Backlog",
	"presenter":{
		"id":1,
		"userId": "admin",
		"name":"admin",
		"roleType":1,
		"lang":"ja",
		"mailAddress":"eguchi@nulab.example"
	},
	"created":"2014-01-23T10:55:19Z"
}`)

	var star Star
	if err := json.Unmarshal(rawJSON, &star); err != nil {
		t.Fatal(err)
	}

	assertInt(t, "ID", star.ID, 75)
	assertString(t, "Comment", star.Comment, "")
	assertString(t, "URL", star.URL, "https://xx.backlog.jp/view/BLG-1")
	assertString(t, "Title", star.Title, "[BLG-1] first issue | 課題の表示 - Backlog")
	assertInt(t, "Presenter#ID", star.Presenter.ID, 1)
	assertTime(t, "Created", star.Created, 2014, 1, 23)
}
