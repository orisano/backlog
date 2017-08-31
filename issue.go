package backlog

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
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

func (c *Client) AddIssue(ctx context.Context, projectID int, summary string, issueTypeID, priorityID int, opt *AddIssueOption) (*AddIssueResponse, error) {
	spath := "/api/v2/issues"
	body := map[string]string{
		"projectId":   fmt.Sprint(projectID),
		"summary":     summary,
		"issueTypeId": fmt.Sprint(issueTypeID),
		"priorityId":  fmt.Sprint(priorityID),
	}

	if opt != nil {
		if len(opt.Description) > 0 {
			body["description"] = opt.Description
		}
	}

	req, err := c.newRequest(ctx, http.MethodPost, spath, &requestOption{
		body: body,
	})
	if err != nil {
		return nil, err
	}

	res, err := c.client.Do(req)
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

func (c *Client) AddIssueComment(ctx context.Context, issueID int, content string) (*AddIssueCommentResponse, error) {
	spath := fmt.Sprintf("/api/v2/issues/%d/comments", issueID)
	req, err := c.newRequest(ctx, http.MethodPost, spath, &requestOption{
		body: map[string]string{
			"content": content,
		},
	})
	if err != nil {
		return nil, err
	}
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err := assertStatusCode(res, http.StatusCreated); err != nil {
		return nil, err
	}

	var out AddIssueCommentResponse
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}
