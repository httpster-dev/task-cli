package task

import "time"

// InMemoryTaskStore holds tasks in a slice — great for testing, no I/O needed.
// Like using a plain Ruby array instead of hitting the database.
type InMemoryTaskStore struct {
	tasks  []Task
	nextID int
}

// NewInMemoryTaskStore is a constructor function — Go's version of .new.
// Returns a pointer so all method calls share the same data.
func NewInMemoryTaskStore() *InMemoryTaskStore {
	return &InMemoryTaskStore{nextID: 1}
}

func (s *InMemoryTaskStore) Add(title string) (Task, error) {
	t := Task{
		ID:        s.nextID,
		Title:     title,
		Status:    StatusTodo,
		Priority:  PriorityMedium,
		CreatedAt: time.Now(),
	}
	s.tasks = append(s.tasks, t)
	s.nextID++
	return t, nil
}

func (s *InMemoryTaskStore) List() ([]Task, error) {
	return s.tasks, nil
}
