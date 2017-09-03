package backlog

import (
	"context"
	"net/http"
)

const (
	diskUsagePath = "/api/v2/space/diskUsage"
)

type DiskUsage struct {
	Capacity   int `json:"capacity"`
	Issue      int `json:"issue"`
	Wiki       int `json:"wiki"`
	File       int `json:"file"`
	Subversion int `json:"subversion"`
	Git        int `json:"git"`
	Details    []struct {
		ProjectID  int `json:"projectId"`
		Issue      int `json:"issue"`
		Wiki       int `json:"wiki"`
		File       int `json:"file"`
		Subversion int `json:"subversion"`
		Git        int `json:"git"`
	} `json:"details"`
}

func (c *Client) GetDiskUsage(ctx context.Context) (*DiskUsage, error) {
	var out DiskUsage
	if err := c.get(ctx, diskUsagePath, http.StatusOK, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
