package backlog

import (
	"context"
	"net/http"
	"path"
)

const (
	usersPath = "/api/v2/users"
)

type User struct {
	ID     int    `json:"id"`
	UserID string `json:"userId"`
	Name   string `json:"name"`
}

func (c *Client) GetMyself(ctx context.Context) (*User, error) {
	spath := path.Join(usersPath, "myself")
	var user User
	if err := c.get(ctx, spath, http.StatusOK, &user); err != nil {
		return nil, err
	}
	return &user, nil
}
