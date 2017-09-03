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

type IssueType struct {
	ID           int    `json:"id"`
	ProjectID    int    `json:"projectId"`
	Name         string `json:"name"`
	Color        string `json:"color"`
	DisplayOrder int    `json:"displayOrder"`
}

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

	var out AddIssueResponse
	if err := c.post(ctx, issuesPath, form, http.StatusCreated, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) AddIssueComment(ctx context.Context, issueID int, content string, optional url.Values) (*AddIssueCommentResponse, error) {
	spath := path.Join(issuesPath, fmt.Sprint(issueID), "comments")

	required := url.Values{}
	required.Set("content", content)

	form := mergeValues(required, optional)

	var out AddIssueCommentResponse
	if err := c.post(ctx, spath, form, http.StatusCreated, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
