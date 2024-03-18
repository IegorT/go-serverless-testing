package entities

import "time"

type Task struct {
	ID          string
	Title       string
	Description string
	OwnerID     string
	Deadline    time.Time
}
