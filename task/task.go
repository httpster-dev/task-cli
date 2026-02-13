package task

import "time"

// Task is the core domain struct — like an AR model but with explicit fields.
// CompletedAt is a pointer so it can be nil (Go's way of expressing "optional").
type Task struct {
	ID          int
	Title       string
	Status      Status
	Priority    Priority
	CreatedAt   time.Time
	CompletedAt *time.Time
}

// Status represents the state of a task — like an enum in Rails.
// Go uses iota to auto-increment integer constants within a const block.
type Status int

const (
	StatusTodo Status = iota
	StatusDone
)

func (s Status) String() string {
	switch s {
	case StatusDone:
		return "done"
	default:
		return "todo"
	}
}

// Priority represents how urgent a task is.
type Priority int

const (
	PriorityLow Priority = iota
	PriorityMedium
	PriorityHigh
)

func (p Priority) String() string {
	switch p {
	case PriorityMedium:
		return "medium"
	case PriorityHigh:
		return "high"
	default:
		return "low"
	}
}
