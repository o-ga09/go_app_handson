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
	ID TaskID `json:"id"`
	Title string `json:"title"`
	Status TaskStatus `json:"status"`
	Created_at time.Time `json:"create_at"`
}

type Tasks []*Task