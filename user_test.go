package backlog

import (
	"encoding/json"
	"testing"
)

func TestParseUser(t *testing.T) {
	rawJSON := []byte(`{
	"id": 1,
    "userId": "admin",
    "name": "admin",
    "roleType": 1,
    "lang": "ja",
    "mailAddress": "eguchi@nulab.example"
}`)
	var user User
	if err := json.Unmarshal(rawJSON, &user); err != nil {
		t.Fatal(err)
	}

	assertInt(t, "ID", user.ID, 1)
	assertString(t, "UserID", user.UserID, "admin")
	assertString(t, "Name", user.Name, "admin")
	assertInt(t, "RoleType", user.RoleType, 1)
	assertString(t, "Lang", user.Lang, "ja")
	assertString(t, "MailAddress", user.MailAddress, "eguchi@nulab.example")
}
