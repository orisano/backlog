package backlog

import (
	"context"
	"encoding/json"
	"net/http"
)

type Priority struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (c *Client) GetPriorities(ctx context.Context) ([]Priority, error) {
	spath := "/api/v2/priorities"
	req, err := c.newRequest(ctx, http.MethodGet, spath, &requestOption{})
	if err != nil {
		return nil, err
	}
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err := assertStatusCode(res, http.StatusOK); err != nil {
		return nil, err
	}

	priorities := make([]Priority, 0)
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&priorities); err != nil {
		return nil, err
	}
	return priorities, nil
}
