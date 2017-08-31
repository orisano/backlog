package backlog

import (
	"context"
	"net/http"
)

const (
	prioritiesPath = "/api/v2/priorities"
)

type Priority struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (c *Client) GetPriorities(ctx context.Context) ([]Priority, error) {
	var priorities []Priority
	if err := c.get(ctx, prioritiesPath, http.StatusOK, &priorities); err != nil {
		return nil, err
	}
	return priorities, nil
}
