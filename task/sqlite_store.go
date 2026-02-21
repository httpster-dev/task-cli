package task

import (
	"database/sql"
	"time"

	_ "modernc.org/sqlite"
)

type SQLiteStore struct {
	db *sql.DB
}

func NewSQLiteStore(path string) (*SQLiteStore, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS tasks (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT,
    status INTEGER,
    priority INTEGER,
    created_at DATETIME,
    completed_at DATETIME
)`)
	if err != nil {
		return nil, err
	}
	return &SQLiteStore{db: db}, nil
}

func (s *SQLiteStore) Add(title string) (Task, error) {
	now := time.Now()
	// INSERT the task into the database
	result, err := s.db.Exec("INSERT INTO tasks (title, status, priority, created_at) VALUES (?, ?, ?, ?)", title, StatusTodo, PriorityMedium, now)
	if err != nil {
		return Task{}, err
	}
	id, err := result.LastInsertId() // this gets the auto-generated ID
	if err != nil {
		return Task{}, err
	}

	// Build and return the Task struct
	t := Task{
		ID:        int(id), // from LastInsertId()
		Title:     title,
		Status:    StatusTodo, // no prefix needed, same package
		Priority:  PriorityMedium,
		CreatedAt: now,
	}
	return t, nil
}

func (s *SQLiteStore) List() ([]Task, error) {
	results, err := s.db.Query("SELECT id, title, status, priority, created_at, completed_at FROM tasks")
	if err != nil {
		return nil, err
	}
	defer results.Close()
	var tasks []Task
	for results.Next() {
		var t Task
		err := results.Scan(&t.ID, &t.Title, &t.Status, &t.Priority, &t.CreatedAt, &t.CompletedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}
