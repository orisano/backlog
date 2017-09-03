package backlog

import (
	"encoding/json"
	"testing"
)

func TestParseActivity(t *testing.T) {
	rawJSON := []byte(`{
	"id": 3153,
	"project": {
		"id": 92,
		"projectKey": "SUB",
		"name": "サブタスク",
		"chartEnabled": true,
		"subtaskingEnabled": true,
		"projectLeaderCanEditProjectLeader": false,
		"textFormattingRule": null,
		"archived": false,
		"displayOrder": 0
	},
	"type": 2,
	"content": {
		"id": 4809,
		"key_id": 121,
		"summary": "コメント",
		"description": "test",
		"comment": {
			"id": 7237,
			"content": "cont"
		},
		"changes": [
			{
				"field": "milestone",
				"new_value": "R2014-07-23",
				"old_value": "a",
				"type": "standard"
			},
			{
				"field": "status",
				"new_value": "4",
				"old_value": "1",
				"type": "standard"
			}
		]
	},
	"notifications": [
		{
			"id": 25,
			"alreadyRead": true,
			"reason": 2,
			"user": {
				"id": 5686,
				"userId": "takada",
				"name": "takada",
				"roleType": 2,
				"lang": "ja",
				"mailAddress": "takada@nulab.example"
			},
			"resourceAlreadyRead":true
		}
	],
	"createdUser": {
		"id": 1,
		"userId": "admin",
		"name": "admin",
		"roleType": 1,
		"lang": "ja",
		"mailAddress": "eguchi@nulab.example"
	},
	"created": "2013-12-27T07:50:44Z"
}`)
	var activity Activity
	if err := json.Unmarshal(rawJSON, &activity); err != nil {
		t.Fatal(err)
	}

	assertInt(t, "ID", activity.ID, 3153)
	assertInt(t, "Project#ID", activity.Project.ID, 92)
	assertInt(t, "Type", activity.Type, 2)
	assertInt(t, "Content#ID", activity.Content.ID, 4809)
	assertInt(t, "Content#KeyID", activity.Content.KeyID, 121)
	assertString(t, "Content#Summary", activity.Content.Summary, "コメント")
	assertString(t, "Content#Description", activity.Content.Description, "test")
	assertInt(t, "Content#Comment#ID", activity.Content.Comment.ID, 7237)
	assertString(t, "Content#Comment#Content", activity.Content.Comment.Content, "cont")
	assertString(t, "Content#Changes[0]#Field", activity.Content.Changes[0].Field, "milestone")
	assertString(t, "Content#Changes[0]#NewValue", activity.Content.Changes[0].NewValue, "R2014-07-23")
	assertString(t, "Content#Changes[0]#OldValue", activity.Content.Changes[0].OldValue, "a")
	assertInt(t, "Notifications[0]#ID", activity.Notifications[0].ID, 25)
	assertBool(t, "Notifications[0]#AlreadyRead", activity.Notifications[0].AlreadyRead, true)
	assertInt(t, "Notifications[0]#Reason", activity.Notifications[0].Reason, 2)
	assertInt(t, "Notifications[0]#User#ID", activity.Notifications[0].User.ID, 5686)
	assertBool(t, "Notifications[0]#ResourceAlreadyRead", activity.Notifications[0].ResourceAlreadyRead, true)
	assertInt(t, "CreatedUser", activity.CreatedUser.ID, 1)
	assertTime(t, "Created", activity.Created, 2013, 12, 27)
}
