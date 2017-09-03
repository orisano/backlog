package backlog

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/pkg/errors"
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

func (c *Client) GetUserReceivedStars(ctx context.Context, userID int) ([]Star, error) {
	spath := path.Join(usersPath, fmt.Sprint(userID), "stars")
	var stars []Star
	if err := c.get(ctx, spath, http.StatusOK, &stars); err != nil {
		return nil, err
	}
	return stars, nil
}

func (c *Client) GetUserReceivedStarsCount(ctx context.Context, userID int, since, until time.Time) (*int, error) {
	spath := path.Join(usersPath, fmt.Sprint(userID), "stars", "count")
	req, err := c.newRequest(ctx, http.MethodGet, spath, nil)
	if err != nil {
		return nil, errors.Wrap(err, "request construct failed")
	}
	query := url.Values{}
	query.Set("since", since.Format(time.RFC3339))
	query.Set("until", until.Format(time.RFC3339))
	req.URL.RawQuery = query.Encode()

	var out struct {
		Count int `json:"count"`
	}
	if err := c.do(req, http.StatusOK, &out); err != nil {
		return nil, errors.Wrap(err, "request failed")
	}
	return &out.Count, nil
}
