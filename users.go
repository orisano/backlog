package backlog

import (
	"context"
	"encoding/json"
	"net/http"
)

type User struct {
	ID     int    `json:"id"`
	UserID string `json:"userId"`
	Name   string `json:"name"`
}

func (c *Client) GetMyself(ctx context.Context) (*User, error) {
	spath := "/api/v2/users/myself"
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

	var user User
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
