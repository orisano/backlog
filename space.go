package backlog

import (
	"context"
	"net/http"
	"time"
)

const (
	spacePath = "/api/v2/space"
)

type Space struct {
	SpaceKey           string    `json:"spaceKey"`
	Name               string    `json:"name"`
	OwnerID            int       `json:"ownerId"`
	Lang               string    `json:"lang"`
	Timezone           string    `json:"timezone"`
	ReportSendTime     string    `json:"reportSendTime"`
	TextFormattingRule string    `json:"textFormattingRule"`
	Created            time.Time `json:"created"`
	Updated            time.Time `json:"updated"`
}

func (c *Client) GetSpace(ctx context.Context) (*Space, error) {
	var out Space
	if err := c.get(ctx, spacePath, http.StatusOK, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
