package main

import (
	"database/sql"
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

func completeTask(db *sql.DB, id int) error {
	query := `
		UPDATE tasks SET completed = 1 WHERE id = ?
	`

	_, err := db.Exec(query, id)
	return err
}

func deleteTask(db *sql.DB, id int) error {
	query := `DELETE FROM tasks WHERE id = ?`
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
