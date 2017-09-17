package backlog

import (
	"context"
	"net/http"
	"path"
)

const (
	categoriesPath = "categories"
)

type Category struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	DisplayOrder int    `json:"displayOrder"`
}

func (c *Client) GetCategories(ctx context.Context, projectID string) ([]Category, error) {
	spath := path.Join(projectsPath, projectID, categoriesPath)
	var categories []Category
	if err := c.get(ctx, spath, http.StatusOK, &categories); err != nil {
		return nil, err
	}
	return categories, nil
}
