package backlog

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"path"
)

const (
	usersPath = "/api/v2/users"
)

const (
	RoleAdmin = 1 + iota
	RoleUser
	RoleReporter
	RoleViewer
	RoleGuestReporter
	RoleGuestViewer
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

func (c *Client) GetUser(ctx context.Context, userID int) (*User, error) {
	spath := path.Join(usersPath, fmt.Sprint(userID))
	var user User
	if err := c.get(ctx, spath, http.StatusOK, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (c *Client) AddUser(ctx context.Context, userID, password, name, mailAddress string, roleType int) (*User, error) {
	params := url.Values{}
	params.Set("userId", userID)
	params.Set("password", password)
	params.Set("name", name)
	params.Set("mailAddress", mailAddress)
	params.Set("roleType", fmt.Sprint(roleType))

	var user User
	if err := c.post(ctx, usersPath, params, http.StatusCreated, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (c *Client) GetMyself(ctx context.Context) (*User, error) {
	spath := path.Join(usersPath, "myself")
	var user User
	if err := c.get(ctx, spath, http.StatusOK, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (c *Client) GetUserActivities(ctx context.Context, userID int) ([]Activity, error) {
	spath := path.Join(usersPath, fmt.Sprint(userID), "activities")
	var activities []Activity
	if err := c.get(ctx, spath, http.StatusOK, &activities); err != nil {
		return nil, err
	}
	return activities, nil
}
