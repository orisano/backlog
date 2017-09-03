package backlog

import "time"

type Notification struct {
	Content string    `json:"content"`
	Updated time.Time `json:"updated"`
}
