package backlog

import (
	"encoding/json"
	"testing"
)

func TestParseCategory(t *testing.T) {
	rawJSON := []byte(`{
	"id": 12,
	"name": "開発",
	"displayOrder": 1
}`)
	var category Category
	if err := json.Unmarshal(rawJSON, &category); err != nil {
		t.Fatal(err)
	}

	assertInt(t, "ID", category.ID, 12)
	assertString(t, "Name", category.Name, "開発")
	assertInt(t, "DisplayOrder", category.DisplayOrder, 1)
}
