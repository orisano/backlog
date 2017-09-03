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
	ID          int    `json:"id"`
	UserID      string `json:"userId"`
	Name        string `json:"name"`
	RoleType    int    `json:"roleType"`
	Lang        string `json:"lang"`
	MailAddress string `json:"mailAddress"`
}

func (c *Client) GetUsers(ctx context.Context) ([]User, error) {
	var users []User
	if err := c.get(ctx, usersPath, http.StatusOK, &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (c *Client) GetMyself(ctx context.Context) (*User, error) {
	spath := path.Join(usersPath, "myself")
	var user User
	if err := c.get(ctx, spath, http.StatusOK, &user); err != nil {
		return nil, err
	}
	return &user, nil
}
