package main

import (
	"database/sql"
	"os"
)

type Task struct {
	ID        int
	Task      string
	Completed bool
}

func createTask(db *sql.DB, content string) error {
	query := `
		INSERT INTO tasks (task) VALUES (?)
	`

	_, err := db.Exec(query, content)

	return err
}

func readTasks(db *sql.DB) ([]Task, error) {
	query := `
		 SELECT id, task, completed from tasks
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	var tasks []Task

	for rows.Next() {
		var t Task
		rows.Scan(&t.ID, &t.Task, &t.Completed)
		tasks = append(tasks, t)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return tasks, nil
}

func openTasks(name string) (*os.File, error) {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		file, err := os.Create(name)
		if err != nil {
			return nil, err
		}
		return file, nil
	}

	return os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0)
}
