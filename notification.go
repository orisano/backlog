package backlog

import (
	"time"
	"context"
	"net/http"
)

const (
	notificationPath = "/api/v2/space/notification"
)

type Notification struct {
	Content string    `json:"content"`
	Updated time.Time `json:"updated"`
}

func (c *Client) GetNotification(ctx context.Context) (*Notification, error) {
	var out Notification
	if err := c.get(ctx, notificationPath, http.StatusOK, &out); err != nil {
		return nil, err
	}
	return &out, nil
}