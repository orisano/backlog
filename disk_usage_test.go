package backlog

import (
	"encoding/json"
	"testing"
)

func TestParseDiskUsage(t *testing.T) {
	rawJSON := []byte(`{
    "capacity": 1073741824,
    "issue": 119511,
    "wiki": 48575,
    "file": 1,
    "subversion": 2,
    "git": 3,
    "details":[
        {
            "projectId": 1,
            "issue": 11931,
            "wiki": 4,
            "file": 5,
            "subversion": 6,
            "git": 7
        }
    ]
}`)
	var diskUsage DiskUsage
	if err := json.Unmarshal(rawJSON, &diskUsage); err != nil {
		t.Fatal(err)
	}

	assertInt(t, "Capacity", diskUsage.Capacity, 1073741824)
	assertInt(t, "Issue", diskUsage.Issue, 119511)
	assertInt(t, "Wiki", diskUsage.Wiki, 48575)
	assertInt(t, "File", diskUsage.File, 1)
	assertInt(t, "Subversion", diskUsage.Subversion, 2)
	assertInt(t, "Git", diskUsage.Git, 3)
	assertInt(t, "Details[0].ProjectID", diskUsage.Details[0].ProjectID, 1)
	assertInt(t, "Details[0].Issue", diskUsage.Details[0].Issue, 11931)
	assertInt(t, "Details[0].Wiki", diskUsage.Details[0].Wiki, 4)
	assertInt(t, "Details[0].File", diskUsage.Details[0].File, 5)
	assertInt(t, "Details[0].Subversion", diskUsage.Details[0].Subversion, 6)
	assertInt(t, "Details[0].Git", diskUsage.Details[0].Git, 7)

}
