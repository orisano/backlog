package backlog

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Project struct {
	ID         int    `json:"id"`
	ProjectKey string `json:"projectKey"`
	Name       string `json:"name"`
}

type IssueType struct {
	ID        int    `json:"id"`
	ProjectID int    `json:"projectId"`
	Name      string `json:"name"`
}

func (c *Client) GetProjects(ctx context.Context) ([]Project, error) {
	spath := "/api/v2/projects"
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

	projects := make([]Project, 0)
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&projects); err != nil {
		return nil, err
	}
	return projects, nil
}

func (c *Client) GetIssueTypes(ctx context.Context, projectID int) ([]IssueType, error) {
	spath := fmt.Sprintf("/api/v2/projects/%d/issueTypes", projectID)
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

	issueTypes := make([]IssueType, 0)
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&issueTypes); err != nil {
		return nil, err
	}
	return issueTypes, nil
}
