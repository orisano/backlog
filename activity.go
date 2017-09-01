package backlog

import (
	"context"
	"net/http"
	"time"
)

const (
	activitiesPath = "/api/v2/space/activities"
)

type Activity struct {
	ID      int     `json:"id"`
	Project Project `json:"project"`
	Type    int     `json:"type"`
	Content struct {
		ID          int    `json:"id"`
		KeyID       int    `json:"key_id"`
		Summary     string `json:"summary"`
		Description string `json:"description"`
		Comment     struct {
			ID      int    `json:"id"`
			Content string `json:"content"`
		}
		Changes []struct {
			Field    string `json:"field"`
			NewValue string `json:"new_value"`
			OldValue string `json:"old_value"`
			Type     string `json:"type"`
		} `json:"changes"`
	} `json:"content"`

	Notifications []struct {
		ID                  int  `json:"id"`
		AlreadyRead         bool `json:"alreadyRead"`
		Reason              int  `json:"reason"`
		User                User `json:"user"`
		ResourceAlreadyRead bool `json:"resourceAlreadyRead"`
	} `json:"notifications"`

	CreatedUser User      `json:"createdUser"`
	Created     time.Time `json:"created"`
}

func (c *Client) GetActivities(ctx context.Context) ([]Activity, error) {
	var activities []Activity
	if err := c.get(ctx, activitiesPath, http.StatusOK, &activities); err != nil {
		return nil, err
	}
	return activities, nil
}
