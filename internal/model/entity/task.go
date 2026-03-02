package entity

import "time"

type Task struct {
	Id          int64      `gorm:"column:id;primary_key"`
	Title       string     `gorm:"column:title"`
	Description *string    `gorm:"column:description"`
	Priority    int        `gorm:"column:priority"`
	DueDate     time.Time  `gorm:"column:due_date"`
	CreatedBy   *string    `gorm:"column:created_by"`
	CreatedAt   *time.Time `gorm:"column:created_at"`
	UpdatedBy   *string    `gorm:"column:updated_by"`
	UpdatedAt   *time.Time `gorm:"column:updated_at"`
}
