package entity

import "time"

type TaskID int64
type TaskStatus string

const (
	TaskStatusTodo TaskStatus = "todo"
	TaskStatusdoing TaskStatus = "doing"
	TaskStatusDone TaskStatus = "done"
)

type Task struct {
	ID TaskID `json:"id" db:"id"`
	UserID UserID `json:"user_id" db:"user_id"`
	Title string `json:"title" db:"title"`
	Status TaskStatus `json:"status" db:"status"`
	Created_at time.Time `json:"create_at" db:"created_at"`
	Modified_at time.Time `json:"modified_at" db:"modified_at"`
}

type Tasks []*Task