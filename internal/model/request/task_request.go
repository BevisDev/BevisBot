package request

import "time"

type Task struct {
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Priority    int       `json:"priority"`
	CreatedBy   string    `json:"created_by"`
}

type TaskSearch struct {
	Q    string `query:"q"`
	Page int    `query:"page"`
	Size int    `query:"size"`
}
