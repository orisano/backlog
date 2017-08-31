package backlog

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"path"
)

const (
	issuesPath = "/api/v2/issues"
)

type AddIssueOption struct {
	Description string
	AssigneeID  int
}

type AddIssueResponse struct {
	IssueID int `json:"id"`
}

type AddIssueCommentResponse struct {
	CommentID int `json:"id"`
}

func (c *Client) AddIssue(ctx context.Context, projectID int, summary string, issueTypeID, priorityID int, optional url.Values) (*AddIssueResponse, error) {
	required := url.Values{}
	required.Set("projectId", fmt.Sprint(projectID))
	required.Set("summary", summary)
	required.Set("issueTypeId", fmt.Sprint(issueTypeID))
	required.Set("priorityId", fmt.Sprint(priorityID))

	form := mergeValues(required, optional)
	body := encodeForm(form)

	var out AddIssueResponse
	if err := c.doSimple(ctx, http.MethodPost, issuesPath, body, http.StatusCreated, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) AddIssueComment(ctx context.Context, issueID int, content string, optional url.Values) (*AddIssueCommentResponse, error) {
	spath := path.Join(issuesPath, fmt.Sprint(issueID), "comments")

	required := url.Values{}
	required.Set("content", content)

	form := mergeValues(required, optional)
	body := encodeForm(form)

	var out AddIssueCommentResponse
	if err := c.doSimple(ctx, http.MethodPost, spath, body, http.StatusCreated, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
