package backlog

import "time"

type Star struct {
	ID        int       `json:"id"`
	Comment   string    `json:"comment"`
	URL       string    `json:"url"`
	Title     string    `json:"title"`
	Presenter User      `json:"presenter"`
	Created   time.Time `json:"created"`
}
