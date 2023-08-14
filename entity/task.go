package entity

type TaskID int
type TaskStatus string

const (
  TaskStatusTodo  TaskStatus = "todo"
  TaskStatusDoing TaskStatus = "doing"
  TaskStatusDone  TaskStatus = "done"
)

type Task struct {
  ID      TaskID     `json:"id"       db:"id"`
  Title   string     `json:"title"    db:"title"`
  Status  TaskStatus `json:"status"   db:"status"`
  Created string     `json:"created"  db:"created"`
  Modified string    `json:"modified" db:"modified"`
}

type Tasks []*Task

