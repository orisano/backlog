package backlog

import (
	"context"
	"fmt"
	"net/http"
	"path"
)

const (
	projectsPath = "/api/v2/projects"
)

type Project struct {
	ID                                int    `json:"id"`
	ProjectKey                        string `json:"projectKey"`
	Name                              string `json:"name"`
	ChartEnabled                      bool   `json:"chartEnabled"`
	SubtaskingEnabled                 bool   `json:"subtaskingEnabled"`
	ProjectLeaderCanEditProjectLeader bool   `json:"projectLeaderCanEditProjectLeader"`
	TextFormattingRule                string `json:"textFormattingRule"`
	Archived                          bool   `json:"archived"`
}

func (c *Client) GetProjects(ctx context.Context) ([]Project, error) {
	var projects []Project
	if err := c.get(ctx, projectsPath, http.StatusOK, &projects); err != nil {
		return nil, err
	}
	return projects, nil
}

func (c *Client) GetIssueTypes(ctx context.Context, projectID int) ([]IssueType, error) {
	spath := path.Join(projectsPath, fmt.Sprint(projectID), "issueTypes")
	var issueTypes []IssueType
	if err := c.get(ctx, spath, http.StatusOK, &issueTypes); err != nil {
		return nil, err
	}
	return issueTypes, nil
}
