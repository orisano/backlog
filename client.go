package backlog

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
	"runtime"
	"strings"
)

const (
	version = "0.1"
)

type Client struct {
	URL        *url.URL
	HTTPClient *http.Client

	APIToken string

	Logger *log.Logger
}

type requestOption struct {
	params map[string]string
	body   map[string]string
}

type Project struct {
	Id         int    `json:"id"`
	ProjectKey string `json:"projectKey"`
	Name       string `json:"name"`
}

type Priority struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type IssueType struct {
	Id        int    `json:"id"`
	ProjectId int    `json:"projectId"`
	Name      string `json:"name"`
}

type User struct {
	Id     int    `json:"id"`
	UserId string `json:"userId"`
	Name   string `json:"name"`
}

type AddIssueOption struct {
	Description string
	AssigneeId  int
}

type AddIssueResponse struct {
	IssueId int `json:"id"`
}

func NewClient(urlStr, apiToken string, logger *log.Logger) (*Client, error) {
	if len(apiToken) == 0 {
		return nil, errors.New("missing token")
	}
	parsedURL, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse url: %s", urlStr)
	}
	var discardLogger = log.New(ioutil.Discard, "", log.LstdFlags)
	if logger == nil {
		logger = discardLogger
	}

	return &Client{
		URL:        parsedURL,
		HTTPClient: http.DefaultClient,

		APIToken: apiToken,

		Logger: logger,
	}, nil
}

var userAgent = fmt.Sprintf("orisano-backlog/%s (%s)", version, runtime.Version())

func (c *Client) newRequest(ctx context.Context, method, spath string, opt *requestOption) (*http.Request, error) {
	if ctx == nil {
		return nil, errors.New("nil context")
	}
	if len(method) == 0 {
		return nil, errors.New("missing method")
	}
	if len(spath) == 0 {
		return nil, errors.New("missing spath")
	}

	u := *c.URL
	u.Path = path.Join(c.URL.Path, spath)

	var r io.Reader
	if len(opt.body) != 0 {
		kv := make([]string, 0, len(opt.body))
		for k, v := range opt.body {
			kv = append(kv, fmt.Sprintf("%s=%s", k, v))
		}
		r = strings.NewReader(strings.Join(kv, "&"))
	}
	req, err := http.NewRequest(method, u.String(), r)
	if err != nil {
		return nil, err
	}

	values := req.URL.Query()
	values.Add("apiKey", c.APIToken)
	if len(opt.params) != 0 {
		for k, v := range opt.params {
			values.Add(k, v)
		}
	}
	req.URL.RawQuery = values.Encode()

	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	req = req.WithContext(ctx)

	return req, nil
}

func (c *Client) AddIssue(ctx context.Context, projectId int, summary string, issueTypeId, priorityId int, opt *AddIssueOption) (*AddIssueResponse, error) {
	spath := "/api/v2/issues"
	body := map[string]string{
		"projectId":   fmt.Sprint(projectId),
		"summary":     summary,
		"issueTypeId": fmt.Sprint(issueTypeId),
		"priorityId":  fmt.Sprint(priorityId),
	}
	req, err := c.newRequest(ctx, "POST", spath, &requestOption{
		body: body,
	})
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err := assertStatusCode(res, http.StatusCreated); err != nil {
		return nil, err
	}

	var out AddIssueResponse
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetProjects(ctx context.Context) ([]Project, error) {
	spath := "/api/v2/projects"
	req, err := c.newRequest(ctx, "GET", spath, &requestOption{})
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err := assertStatusCode(res, http.StatusOK); err != nil {
		return nil, err
	}

	projects := make([]Project, 0)
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&projects); err != nil {
		return nil, err
	}
	return projects, nil
}

func (c *Client) GetPriorities(ctx context.Context) ([]Priority, error) {
	spath := "/api/v2/priorities"
	req, err := c.newRequest(ctx, "GET", spath, &requestOption{})
	if err != nil {
		return nil, err
	}
	res, err := c.HTTPClient.Do(req)
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

func (c *Client) GetIssueTypes(ctx context.Context, projectId int) ([]IssueType, error) {
	spath := fmt.Sprintf("/api/v2/projects/%d/issueTypes", projectId)
	req, err := c.newRequest(ctx, "GET", spath, &requestOption{})
	if err != nil {
		return nil, err
	}
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err := assertStatusCode(res, http.StatusOK); err != nil {
		return nil, err
	}

	issueTypes := make([]IssueType, 0)
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&issueTypes); err != nil {
		return nil, err
	}
	return issueTypes, nil
}

func (c *Client) GetMyself(ctx context.Context) (*User, error) {
	spath := "/api/v2/users/myself"
	req, err := c.newRequest(ctx, "GET", spath, &requestOption{})
	if err != nil {
		return nil, err
	}
	res, err := c.HTTPClient.Do(req)
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

func assertStatusCode(res *http.Response, expected int) error {
	if res.StatusCode != expected {
		return errors.Errorf("invalid status code: %s", res.Status)
	} else {
		return nil
	}
}
