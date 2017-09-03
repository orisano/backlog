package backlog

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
