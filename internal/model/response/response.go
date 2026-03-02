package response

import "time"

type Task struct {
	Id          int64      `json:"id"`
	Title       string     `json:"title"`
	Description *string    `json:"description"`
	Priority    int        `json:"priority"`
	DueDate     time.Time  `json:"due_date"`
	CreatedBy   *string    `json:"created_by"`
	CreatedAt   *time.Time `json:"created_at"`
}
