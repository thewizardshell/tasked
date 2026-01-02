package domain

import "time"

type Task struct {
	Id          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Priority    string    `json:"priority"`
	Userid      int64     `json:"userId"`
	Duedate     time.Time `json:"dueDate"`
	CompletedAt time.Time `json:"completedAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
